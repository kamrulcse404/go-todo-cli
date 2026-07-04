package commands

import "todo-cli/models"

func GetNextID(todos []models.Todo) int {
	maxID := 0

	for _, todo := range todos {
		if todo.ID > maxID {
			maxID = todo.ID
		}
	}

	return maxID + 1
}
