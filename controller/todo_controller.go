package controller

import (
	"net/http"
	"strconv"
	"todo-api-go/model"
	"todo-api-go/usecase"

	"github.com/gin-gonic/gin"
)

type todoController struct {
	todoUseCase usecase.TodoUseCase
}

func NewTodoController(todoUseCase usecase.TodoUseCase) todoController {
	return todoController{
		todoUseCase: todoUseCase,
	}
}

func (tc *todoController) GetTodos(context *gin.Context) {
	todos, err := tc.todoUseCase.GetTodos()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, todos)
}

func (tc *todoController) GetTodo(context *gin.Context) {
	todoId := context.Param("id")
	if todoId == "" {
		response := model.Response{
			Message: "Todo id cannot be empty",
		}
		context.JSON(http.StatusBadRequest, response)
		return
	}

	id, err := strconv.Atoi(todoId)
	if err != nil {
		response := model.Response{
			Message: "Todo id must be a number",
		}
		context.JSON(http.StatusBadRequest, response)
		return
	}

	todo, err := tc.todoUseCase.GetTodo(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	if todo == nil {
		response := model.Response{
			Message: "Todo not found in the database",
		}
		context.JSON(http.StatusNotFound, response)
		return
	}

	context.JSON(http.StatusOK, todo)
}

func (tc *todoController) CreateTodo(context *gin.Context) {
	var body struct {
		Content string `json:"content"`
	}

	err := context.BindJSON(&body)
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	if body.Content == "" {
		response := model.Response{
			Message: "Missing param: content",
		}
		context.JSON(http.StatusBadRequest, response)
		return
	}

	if len(body.Content) > 255 {
		response := model.Response{
			Message: "Content has more than 255 characters",
		}
		context.JSON(http.StatusBadRequest, response)
		return
	}

	createdTodo, err := tc.todoUseCase.CreateTodo(body.Content)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusCreated, createdTodo)
}

func (tc *todoController) ChangeTodoStatus(context *gin.Context) {
	todoId := context.Param("id")
	if todoId == "" {
		response := model.Response{
			Message: "Todo id cannot be empty",
		}
		context.JSON(http.StatusBadRequest, response)
		return
	}

	id, err := strconv.Atoi(todoId)
	if err != nil {
		response := model.Response{
			Message: "Todo id must be a number",
		}
		context.JSON(http.StatusBadRequest, response)
		return
	}

	err = tc.todoUseCase.ChangeTodoStatus(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusNoContent, nil)
}

func (tc *todoController) DeleteTodo(context *gin.Context) {
	todoId := context.Param("id")
	if todoId == "" {
		response := model.Response{
			Message: "Todo id cannot be empty",
		}
		context.JSON(http.StatusBadRequest, response)
		return
	}

	id, err := strconv.Atoi(todoId)
	if err != nil {
		response := model.Response{
			Message: "Todo id must be a number",
		}
		context.JSON(http.StatusBadRequest, response)
		return
	}

	err = tc.todoUseCase.DeleteTodo(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusNoContent, nil)
}
