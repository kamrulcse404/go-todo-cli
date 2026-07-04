package commands

import (
	"fmt"
	"todo-cli/storage"
)

func CompleteTodo(id int) {
	if id <= 0 {
		fmt.Println("Invalid ID!")
		return
	}

	todos, err := storage.LoadTodos()
	if err != nil {
		fmt.Println("Error loading todos:", err)
		return
	}

	found := false

	for idx := range todos {
		if todos[idx].ID == id {
			if todos[idx].Completed {
				fmt.Println("Todo is already completed!")
				return
			}
			todos[idx].Completed = true
			found = true
			break

		}
	}

	if !found {
		fmt.Println("Todo not found!")
		return
	}

	err = storage.SaveTodos(todos)
	if err != nil {
		fmt.Println("Error saving todos:", err)
		return
	}
	fmt.Printf("Todo %d marked as completed!\n", id)

}
