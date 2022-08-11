package controllers

import (
	"example/todo-go/initializers"
	"example/todo-go/middleware"
	"example/todo-go/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexPage(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{
		"content": "This is an index page...",
	})
}

func CreateTodo(context *gin.Context) {

	var body struct {
		Item string
	}

	context.Bind(&body)

	todo := models.Todo{Item: body.Item, Completed: false, UserID: middleware.Id}
	result := initializers.DB.Create(&todo)

	if result.Error != nil {
		context.Status(400)
		return
	}

	context.JSON(200, gin.H{
		"todo": todo,
	})
}

func GetTodoList(context *gin.Context) {
	var todoList []models.Todo

	initializers.DB.Find(&todoList)

	context.JSON(200, gin.H{
		"todoList": todoList,
	})
}

func GetTodo(context *gin.Context) {

	id := context.Param("id")

	var todo models.Todo

	initializers.DB.First(&todo, id)

	context.JSON(200, gin.H{
		"todo": todo,
	})
}

func UpdateTodo(context *gin.Context) {
	id := context.Param("id")

	var body struct {
		Item      string
		Completed bool
	}

	context.Bind(&body)
	var todo models.Todo

	initializers.DB.First(&todo, id)
	initializers.DB.Model(&todo).Updates(models.Todo{
		Item: body.Item, Completed: body.Completed,
	})
	context.JSON(200, gin.H{
		"todo": todo,
	})
}

func DeleteTodo(context *gin.Context) {
	id := context.Param("id")

	initializers.DB.Delete(&models.Todo{}, id)

	context.Status(200)
}
