package repository

import (
	"database/sql"
	"fmt"
	"todo-api-go/model"
)

type TodoRepository struct {
	connection *sql.DB
}

func NewTodoRepository(connection *sql.DB) TodoRepository {
	return TodoRepository{
		connection: connection,
	}
}

func (tr *TodoRepository) GetTodos() ([]model.Todo, error) {
	query := "SELECT * FROM todos"
	rows, err := tr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Todo{}, err
	}

	todoList := []model.Todo{}
	var todo model.Todo

	for rows.Next() {
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.Done,
			&todo.CreatedAt,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Todo{}, err
		}

		todoList = append(todoList, todo)
	}

	rows.Close()

	return todoList, nil
}

func (tr *TodoRepository) GetTodo(id int) (*model.Todo, error) {
	query, err := tr.connection.Prepare("SELECT * FROM todos WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var todo model.Todo

	err = query.QueryRow(id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.Done,
		&todo.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()

	return &todo, nil
}

func (tr *TodoRepository) CreateTodo(content string) (*model.Todo, error) {
	query, err := tr.connection.Prepare("INSERT INTO todos(content) " +
		"VALUES ($1) RETURNING *",
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var createdTodo model.Todo
	err = query.QueryRow(content).Scan(
		&createdTodo.ID,
		&createdTodo.Content,
		&createdTodo.Done,
		&createdTodo.CreatedAt,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	query.Close()

	return &createdTodo, nil
}

func (tr *TodoRepository) ChangeTodoStatus(id int) error {
	query, err := tr.connection.Prepare("UPDATE todos SET done = NOT done WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = query.Exec(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (tr *TodoRepository) DeleteTodo(id int) error {
	query, err := tr.connection.Prepare("DELETE FROM todos WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = query.Exec(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
