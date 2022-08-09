package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todoList = []todo{
	{ID: "1", Item: "Create REST API in go", Completed: false},
	{ID: "2", Item: "Take a shower", Completed: false},
	{ID: "3", Item: "Cook pasta", Completed: false},
}

func getTodoList(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todoList)
}

func main() {
	router := gin.Default()
	router.GET("/todoList", getTodoList)
	router.Run("localhost:8080")
}
