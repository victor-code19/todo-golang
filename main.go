package main

import (
	"context"
	"log"
	"os"
	"time"
	"example.com/todo-rest-api/controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func main() {

	router := gin.Default()
	uc := controllers.NewTaskController(getClient())

	router.Static("/static", "./public")
	router.LoadHTMLGlob("templates/*.gohtml")

	apiRoutes := router.Group("/api")
	viewRoutes := router.Group("/view")

	apiRoutes.POST("/task", uc.CreateTask)
	apiRoutes.GET("/tasks", uc.GetTasks)
	apiRoutes.DELETE("/task/:id", uc.DeleteTask)
	apiRoutes.DELETE("/tasks", uc.DeleteAllTasks)

	viewRoutes.GET("/tasks", uc.ShowAllTasks)

	router.Run(":8080")
}

func getClient() *mongo.Client {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable")
	}
	
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// Sprawdź połączenie
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}
	
	log.Println("Successfully connected to MongoDB!")
	return client
}
