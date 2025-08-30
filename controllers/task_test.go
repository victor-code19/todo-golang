package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"example.com/todo-rest-api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type TaskControllerTestSuite struct {
	suite.Suite
	controller *TaskController
	client     *mongo.Client
	collection *mongo.Collection
}

func (suite *TaskControllerTestSuite) SetupSuite() {
	// Connect to test MongoDB
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		suite.T().Fatalf("Failed to connect to MongoDB: %v", err)
	}

	suite.client = client
	suite.collection = client.Database("todo-app-go-test").Collection("tasks")
	suite.controller = NewTaskControllerWithDB(client, "todo-app-go-test")

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		suite.T().Fatalf("Failed to ping MongoDB: %v", err)
	}
}

func (suite *TaskControllerTestSuite) TearDownSuite() {
	if suite.client != nil {
		suite.client.Disconnect(context.Background())
	}
}

func (suite *TaskControllerTestSuite) SetupTest() {
	// Clean up collection before each test
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	suite.collection.Drop(ctx)
}

func (suite *TaskControllerTestSuite) TestCreateTask() {
	gin.SetMode(gin.TestMode)
	
	taskData := map[string]string{
		"description": "Test task",
	}
	
	jsonData, _ := json.Marshal(taskData)
	
	req, _ := http.NewRequest("POST", "/api/task", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	router := gin.New()
	router.POST("/api/task", suite.controller.CreateTask)
	router.ServeHTTP(w, req)
	
	assert.Equal(suite.T(), http.StatusCreated, w.Code)
	
	var response models.Task
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Test task", response.Description)
	assert.False(suite.T(), response.Id.IsZero())
}

func (suite *TaskControllerTestSuite) TestCreateTaskInvalidJSON() {
	gin.SetMode(gin.TestMode)
	
	req, _ := http.NewRequest("POST", "/api/task", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	router := gin.New()
	router.POST("/api/task", suite.controller.CreateTask)
	router.ServeHTTP(w, req)
	
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Invalid JSON format", response["message"])
}

func (suite *TaskControllerTestSuite) TestGetTasks() {
	gin.SetMode(gin.TestMode)
	
	// Insert test data
	testTasks := []interface{}{
		models.Task{Id: bson.NewObjectID(), Description: "Task 1"},
		models.Task{Id: bson.NewObjectID(), Description: "Task 2"},
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	_, err := suite.collection.InsertMany(ctx, testTasks)
	assert.NoError(suite.T(), err)
	
	req, _ := http.NewRequest("GET", "/api/tasks", nil)
	w := httptest.NewRecorder()
	router := gin.New()
	router.GET("/api/tasks", suite.controller.GetTasks)
	router.ServeHTTP(w, req)
	
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	
	var response []models.Task
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), response, 2)
}

func (suite *TaskControllerTestSuite) TestDeleteTask() {
	gin.SetMode(gin.TestMode)
	
	// Insert test task
	testTask := models.Task{Id: bson.NewObjectID(), Description: "Task to delete"}
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	_, err := suite.collection.InsertOne(ctx, testTask)
	assert.NoError(suite.T(), err)
	
	req, _ := http.NewRequest("DELETE", "/api/task/"+testTask.Id.Hex(), nil)
	w := httptest.NewRecorder()
	router := gin.New()
	router.DELETE("/api/task/:id", suite.controller.DeleteTask)
	router.ServeHTTP(w, req)
	
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	
	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Task deleted successfully", response["message"])
}

func (suite *TaskControllerTestSuite) TestDeleteTaskInvalidID() {
	gin.SetMode(gin.TestMode)
	
	req, _ := http.NewRequest("DELETE", "/api/task/invalid-id", nil)
	w := httptest.NewRecorder()
	router := gin.New()
	router.DELETE("/api/task/:id", suite.controller.DeleteTask)
	router.ServeHTTP(w, req)
	
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
	
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Invalid ID format", response["message"])
}

func (suite *TaskControllerTestSuite) TestDeleteTaskNotFound() {
	gin.SetMode(gin.TestMode)
	
	nonExistentID := bson.NewObjectID()
	req, _ := http.NewRequest("DELETE", "/api/task/"+nonExistentID.Hex(), nil)
	w := httptest.NewRecorder()
	router := gin.New()
	router.DELETE("/api/task/:id", suite.controller.DeleteTask)
	router.ServeHTTP(w, req)
	
	assert.Equal(suite.T(), http.StatusNotFound, w.Code)
	
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Task not found", response["message"])
}

func (suite *TaskControllerTestSuite) TestDeleteAllTasks() {
	gin.SetMode(gin.TestMode)
	
	// Insert test data
	testTasks := []interface{}{
		models.Task{Id: bson.NewObjectID(), Description: "Task 1"},
		models.Task{Id: bson.NewObjectID(), Description: "Task 2"},
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	_, err := suite.collection.InsertMany(ctx, testTasks)
	assert.NoError(suite.T(), err)
	
	req, _ := http.NewRequest("DELETE", "/api/tasks", nil)
	w := httptest.NewRecorder()
	router := gin.New()
	router.DELETE("/api/tasks", suite.controller.DeleteAllTasks)
	router.ServeHTTP(w, req)
	
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "All tasks deleted successfully", response["message"])
	assert.Equal(suite.T(), float64(2), response["deletedCount"])
}

func TestTaskControllerSuite(t *testing.T) {
	suite.Run(t, new(TaskControllerTestSuite))
}
