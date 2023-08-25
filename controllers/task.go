package controllers

import (
	"net/http"

	"example.com/todo-rest-api/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TaskController struct {
	session *mgo.Session
}

func NewTaskController(s *mgo.Session) *TaskController {
	return &TaskController{s}
}

func (uc TaskController) GetTasks(c *gin.Context) {
	result := []models.Task{}

	if err := uc.session.DB("todo-app-go").C("tasks").Find(bson.M{}).All(&result); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Unable to fetch tasks"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (uc TaskController) CreateTask(c *gin.Context) {
	newTask := models.Task{}

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	newTask.Id = bson.NewObjectId()

	if err := uc.session.DB("todo-app-go").C("tasks").Insert(newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	}

	c.JSON(http.StatusCreated, newTask)

}

func (uc TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	if !bson.IsObjectIdHex(id) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "incorrect id"})
		return
	}

	_id := bson.ObjectIdHex(id)

	if err := uc.session.DB("todo-app-go").C("tasks").RemoveId(_id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task deleted"})

}

func (uc TaskController) ShowAllTasks(c *gin.Context) {
	tasks := []models.Task{}

	if err := uc.session.DB("todo-app-go").C("tasks").Find(bson.M{}).All(&tasks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unable to fetch tasks"})
		return
	}

	viewTasks := []models.ViewTask{}

	for _, element := range tasks {
		viewTasks = append(viewTasks, models.ViewTask{element.Id.Hex(), element.Description})
	}

	data := gin.H{
		"tasks":        viewTasks,
		"tasksCounter": len(tasks),
	}

	c.HTML(http.StatusOK, "index.gohtml", data)

}

func (uc TaskController) DeleteAllTasks(c *gin.Context) {
	if _, err := uc.session.DB("todo-app-go").C("tasks").RemoveAll(bson.M{}); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "tasks not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All tasks deleted"})

}
