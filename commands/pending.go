package commands

import (
	"fmt"
	"todo-cli/storage"
)

func PeningTodosList() {
	todos, err := storage.LoadTodos()

	if err != nil {
		fmt.Println("Error loading todos:", err)
		return
	}

	if len(todos) == 0 {
		fmt.Println("No todos found.")
		return
	}

	status := "✗"
	for _, todo := range todos {
		if !todo.Completed {
			fmt.Printf("[%s] %d. %s\n", status, todo.ID, todo.Title)			
		}
	}
}

