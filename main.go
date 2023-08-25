package main

import (
	"example.com/todo-rest-api/controllers"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func main() {

	router := gin.Default()
	uc := controllers.NewTaskController(getSession())

	router.Static("/static", "./static")
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

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		panic(err)
	}

	return s
}
