package main

import (
	"example/todo-go/initializers"
	"example/todo-go/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Todo{})
}
