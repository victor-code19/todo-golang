package controllers

import (
	"context"
	"time"
	"net/http"
    "log"

	"example.com/todo-rest-api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	defaultTimeout = 5 * time.Second
	dbName         = "todo-app-go"
	collectionName = "tasks"
)

type TaskController struct {
	collection *mongo.Collection
}

func NewTaskController(c *mongo.Client) *TaskController {
	return &TaskController{
		collection: c.Database(dbName).Collection(collectionName),
	}
}

func NewTaskControllerWithDB(c *mongo.Client, database string) *TaskController {
	return &TaskController{
		collection: c.Database(database).Collection(collectionName),
	}
}

func (tc TaskController) getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), defaultTimeout)
}

func (tc TaskController) GetTasks(c *gin.Context) {
	ctx, cancel := tc.getContext()
	defer cancel()

	cursor, err := tc.collection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch tasks"})
		return
	}
	defer cursor.Close(ctx)

	var tasks []models.Task

	if err = cursor.All(ctx, &tasks); err != nil {
        log.Println("Error fetching tasks:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error decoding tasks"})
        return
	}

	c.JSON(http.StatusOK, tasks)
}

func (tc TaskController) CreateTask(c *gin.Context) {
	ctx, cancel := tc.getContext()
	defer cancel()

	var newTask models.Task

	if err := c.BindJSON(&newTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format"})
        return
    }

	result, err := tc.collection.InsertOne(ctx, newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create task"})
        return
	}
	
	if oid, ok := result.InsertedID.(bson.ObjectID); ok {
        newTask.Id = oid
    }
	
	c.JSON(http.StatusCreated, newTask)

}

func (tc TaskController) DeleteTask(c *gin.Context) {
	ctx, cancel := tc.getContext()
	defer cancel()

	id := c.Param("id")

	objectID, err := bson.ObjectIDFromHex(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
        return
    }

	result, err := tc.collection.DeleteOne(ctx, bson.M{"_id": objectID})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete task"})
        return
    }

	if result.DeletedCount == 0 {
        c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})

}

func (tc TaskController) ShowAllTasks(c *gin.Context) {
    ctx, cancel := tc.getContext()
    defer cancel()

    cursor, err := tc.collection.Find(ctx, bson.M{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch tasks"})
        return
    }
    defer cursor.Close(ctx)

    var tasks []models.Task
    if err = cursor.All(ctx, &tasks); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Error decoding tasks"})
        return
    }

    viewTasks := make([]models.ViewTask, 0, len(tasks))
    for _, task := range tasks {
        viewTasks = append(viewTasks, models.ViewTask{
            Id:          task.Id.Hex(),
            Description: task.Description,
        })
    }

    data := gin.H{
        "tasks":        viewTasks,
        "tasksCounter": len(tasks),
    }

    c.HTML(http.StatusOK, "index.gohtml", data)
}

func (tc TaskController) DeleteAllTasks(c *gin.Context) {
    ctx, cancel := tc.getContext()
    defer cancel()

    result, err := tc.collection.DeleteMany(ctx, bson.M{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete tasks"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message":      "All tasks deleted successfully",
        "deletedCount": result.DeletedCount,
    })
}