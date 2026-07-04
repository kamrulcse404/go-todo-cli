package storage

import (
	"encoding/json"
	"os"
	"todo-cli/models"
)

func LoadTodos() ([]models.Todo, error) {
	data, err := os.ReadFile("data/todos.json")

	if os.IsNotExist(err) {
		return []models.Todo{}, nil
	}

	if err != nil {
		return nil, err
	}

	var todos []models.Todo
	err = json.Unmarshal(data, &todos)

	if err != nil {
		return nil, err
	}

	return todos, nil
}

func SaveTodos(todos []models.Todo) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("data/todos.json", data, 0644)
	if err != nil {
		return err
	}
	return nil
}