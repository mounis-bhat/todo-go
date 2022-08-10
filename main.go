package main

import (
	"example/todo-go/controllers"
	"example/todo-go/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", controllers.IndexPage)
	router.GET("/todoList", controllers.GetTodoList)
	router.GET("/todoList/:id", controllers.GetTodo)
	router.PUT("/todoList/:id", controllers.UpdateTodo)
	router.POST("/todoList", controllers.CreateTodo)
	router.DELETE("/todoList/:id", controllers.DeleteTodo)

	router.Run()
}
