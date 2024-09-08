package usecase

import (
	"todo-api-go/model"
	"todo-api-go/repository"
)

type TodoUseCase struct {
	repository repository.TodoRepository
}

func NewTodoUseCase(repo repository.TodoRepository) TodoUseCase {
	return TodoUseCase{
		repository: repo,
	}
}

func (tu *TodoUseCase) GetTodos() ([]model.Todo, error) {
	return tu.repository.GetTodos()
}

func (tu *TodoUseCase) GetTodo(id int) (*model.Todo, error) {
	return tu.repository.GetTodo(id)
}

func (tu *TodoUseCase) CreateTodo(content string) (*model.Todo, error) {
	return tu.repository.CreateTodo(content)
}

func (tu *TodoUseCase) ChangeTodoStatus(id int) error {
	return tu.repository.ChangeTodoStatus(id)
}

func (tu *TodoUseCase) DeleteTodo(id int) error {
	return tu.repository.DeleteTodo(id)
}
