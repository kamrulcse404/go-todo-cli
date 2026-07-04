package commands

import (
	"errors"
	"todo-cli/models"
	"todo-cli/storage"
)

func SearchTodo(id int) (models.Todo, error) {
	if id <= 0 {
		return models.Todo{}, errors.New("invalid id")
	}

	todos, err := storage.LoadTodos()
	if err != nil {
		return models.Todo{}, err
	}

	for _, todo := range todos {
		if todo.ID == id {
			return todo, nil
		}
	}

	return models.Todo{}, errors.New("todo not found")
}
