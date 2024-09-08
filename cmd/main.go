package main

import (
	"todo-api-go/controller"
	"todo-api-go/database"
	"todo-api-go/repository"
	"todo-api-go/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	TodoRepository := repository.NewTodoRepository(dbConnection)
	TodoUseCase := usecase.NewTodoUseCase(TodoRepository)
	TodoController := controller.NewTodoController(TodoUseCase)

	server.GET("/todos", TodoController.GetTodos)
	server.GET("/todos/:id", TodoController.GetTodo)
	server.POST("/todos", TodoController.CreateTodo)
	server.PATCH("/todos/:id/status", TodoController.ChangeTodoStatus)
	server.DELETE("/todos/:id", TodoController.DeleteTodo)

	server.Run()
}
