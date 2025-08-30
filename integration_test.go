package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"example.com/todo-rest-api/controllers"
	"example.com/todo-rest-api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type IntegrationTestSuite struct {
	suite.Suite
	router *gin.Engine
	client *mongo.Client
}

func (suite *IntegrationTestSuite) SetupSuite() {
	// Set test environment
	os.Setenv("MONGODB_URI", "mongodb://localhost:27017")
	gin.SetMode(gin.TestMode)

	// Connect to test MongoDB
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		suite.T().Fatalf("Failed to connect to MongoDB: %v", err)
	}

	suite.client = client

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		suite.T().Fatalf("Failed to ping MongoDB: %v", err)
	}

	// Setup router
	suite.router = gin.New()
	uc := controllers.NewTaskControllerWithDB(client, "todo-app-go-test")

	apiRoutes := suite.router.Group("/api")
	viewRoutes := suite.router.Group("/view")

	apiRoutes.POST("/task", uc.CreateTask)
	apiRoutes.GET("/tasks", uc.GetTasks)
	apiRoutes.DELETE("/task/:id", uc.DeleteTask)
	apiRoutes.DELETE("/tasks", uc.DeleteAllTasks)

	viewRoutes.GET("/tasks", uc.ShowAllTasks)
}

func (suite *IntegrationTestSuite) TearDownSuite() {
	if suite.client != nil {
		suite.client.Disconnect(context.Background())
	}
}

func (suite *IntegrationTestSuite) SetupTest() {
	// Clean up test database before each test
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	collection := suite.client.Database("todo-app-go-test").Collection("tasks")
	collection.Drop(ctx)
}

func (suite *IntegrationTestSuite) TestFullTaskWorkflow() {
	// 1. Create a task
	taskData := map[string]string{
		"description": "Integration test task",
	}
	jsonData, _ := json.Marshal(taskData)

	req, _ := http.NewRequest("POST", "/api/task", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusCreated, w.Code)

	var createdTask models.Task
	err := json.Unmarshal(w.Body.Bytes(), &createdTask)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Integration test task", createdTask.Description)
	assert.False(suite.T(), createdTask.Id.IsZero())

	// 2. Get all tasks
	req, _ = http.NewRequest("GET", "/api/tasks", nil)
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)

	var tasks []models.Task
	err = json.Unmarshal(w.Body.Bytes(), &tasks)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), tasks, 1)
	assert.Equal(suite.T(), createdTask.Id, tasks[0].Id)

	// 3. Delete the task
	req, _ = http.NewRequest("DELETE", "/api/task/"+createdTask.Id.Hex(), nil)
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)

	// 4. Verify task is deleted
	req, _ = http.NewRequest("GET", "/api/tasks", nil)
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)

	var emptyTasks []models.Task
	err = json.Unmarshal(w.Body.Bytes(), &emptyTasks)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), emptyTasks, 0)
}

func (suite *IntegrationTestSuite) TestCreateMultipleTasksAndDeleteAll() {
	// Create multiple tasks
	for i := 0; i < 3; i++ {
		taskData := map[string]string{
			"description": "Task " + string(rune('A'+i)),
		}
		jsonData, _ := json.Marshal(taskData)

		req, _ := http.NewRequest("POST", "/api/task", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)

		assert.Equal(suite.T(), http.StatusCreated, w.Code)
	}

	// Verify all tasks exist
	req, _ := http.NewRequest("GET", "/api/tasks", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)

	var tasks []models.Task
	err := json.Unmarshal(w.Body.Bytes(), &tasks)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), tasks, 3)

	// Delete all tasks
	req, _ = http.NewRequest("DELETE", "/api/tasks", nil)
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "All tasks deleted successfully", response["message"])
	assert.Equal(suite.T(), float64(3), response["deletedCount"])

	// Verify all tasks are deleted
	req, _ = http.NewRequest("GET", "/api/tasks", nil)
	w = httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)

	var emptyTasks []models.Task
	err = json.Unmarshal(w.Body.Bytes(), &emptyTasks)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), emptyTasks, 0)
}

func TestIntegrationSuite(t *testing.T) {
	// Skip integration tests if MongoDB is not available
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}

	suite.Run(t, new(IntegrationTestSuite))
}
