package main

import (
	"example/todo-go/controllers"
	"example/todo-go/initializers"
	"example/todo-go/middleware"

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

	router.GET("/todoList", middleware.RequireAuth, controllers.GetTodoList)
	router.GET("/todoList/:id", middleware.RequireAuth, controllers.GetTodo)
	router.PUT("/todoList/:id", middleware.RequireAuth, controllers.UpdateTodo)
	router.POST("/todoList", middleware.RequireAuth, controllers.CreateTodo)
	router.DELETE("/todoList/:id", middleware.RequireAuth, controllers.DeleteTodo)

	router.POST("/signUp", controllers.SignUp)
	router.POST("/login", controllers.Login)

	router.Run()
}
