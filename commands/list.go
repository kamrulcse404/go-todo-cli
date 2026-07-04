package commands

import (
	"fmt"
	"todo-cli/storage"
)

func ListTodos() {
	todos, err := storage.LoadTodos()

	if err != nil {
		fmt.Println("Error loading todos:", err)
		return
	}

	if len(todos) == 0 {
		fmt.Println("No todos found.")
		return
	}

	for _, todo := range todos {
		status := "✗"
		if todo.Completed {
			status = "✓"
		}
		fmt.Printf("[%s] %d. %s\n", status, todo.ID, todo.Title)
	}
}

