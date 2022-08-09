package main

import (
	"errors"
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

func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todoList = append(todoList, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoById(context *gin.Context) {
	id := context.Param("id")

	todo, error := searchTodo(id)

	if error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")

	todo, error := searchTodo(id)

	if error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

func deleteTodo(context *gin.Context) {
	id := context.Param("id")

	found := false

	for index, todo := range todoList {
		if todo.ID == id {
			todoList = append(todoList[:index], todoList[index+1:]...)
			found = true
			break
		}
	}

	if found {
		context.IndentedJSON(http.StatusOK, gin.H{"message": "Todo deleted"})
	} else {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
	}

}

func indexPage(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{
		"content": "This is an index page...",
	})
}

func searchTodo(id string) (*todo, error) {
	// for i := 0; i < len(todoList); i++ {
	// 	if todoList[i].ID == id {
	// 		return &todoList[i], nil
	// 	}
	// }

	for index, todo := range todoList {
		if todo.ID == id {
			return &todoList[index], nil
		}
	}

	return nil, errors.New("todo not found")
}

func main() {
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", indexPage)

	router.GET("/todoList", getTodoList)
	router.GET("/todoList/:id", getTodoById)
	router.PATCH("/todoList/:id", toggleTodoStatus)
	router.POST("/todoList", addTodo)
	router.DELETE("/todoList/:id", deleteTodo)
	router.Run("localhost:8080")
}
