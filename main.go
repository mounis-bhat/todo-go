package main

import (
	"example/todo-go/controllers"
	"example/todo-go/initializers"
	"example/todo-go/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()

	config.AllowOrigins = []string{"https://react-playground-lake-eight.vercel.app", "http://127.0.0.1:5173", "http://localhost:5173"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	config.AllowCredentials = true
	config.ExposeHeaders = []string{"set-cookie"}
	router.Use(cors.New(config))

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", controllers.IndexPage)

	router.GET("/todo-list", middleware.RequireAuth, controllers.GetTodoList)
	router.GET("/todo-list/:id", middleware.RequireAuth, controllers.GetTodo)
	router.PUT("/todo-list/:id", middleware.RequireAuth, controllers.UpdateTodo)
	router.POST("/todo-list", middleware.RequireAuth, controllers.CreateTodo)
	router.DELETE("/todo-list/:id", middleware.RequireAuth, controllers.DeleteTodo)

	router.POST("/sign-up", controllers.SignUp)
	router.POST("/login", controllers.Login)

	router.Run()
}
