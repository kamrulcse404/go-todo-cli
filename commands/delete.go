package commands

import (
	"fmt"
	"todo-cli/storage"
)

func DeleteTodo(id int) {
	if id <= 0 {
		fmt.Println("Invalid ID!")
		return
	}

	todos, err := storage.LoadTodos()
	if err != nil {
		fmt.Println("Error loading todos:", err)
		return
	}

	deleted := false

	for idx := range todos {
		if todos[idx].ID == id {
			todos = append(todos[:idx], todos[idx+1:]...)
			deleted = true
			break
		}
	}

	if !deleted {
		fmt.Println("Todo not found!")
		return
	}

	err = storage.SaveTodos(todos)
	if err != nil {
		fmt.Println("Error saving todos:", err)
		return
	}

	fmt.Printf("Todo %d deleted successfully!\n", id)
}
