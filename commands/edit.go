package commands

import (
	"fmt"
	"todo-cli/storage"
)

func EditTodo(id int, title string) {

	if id <= 0 {
		fmt.Println("Invalid ID!")
		return
	}

	if title == "" {
		fmt.Println("Todo title cannot be empty.")
		return
	}

	todos, err := storage.LoadTodos()
	if err != nil {
		fmt.Println("Error loading todos!", err)
		return
	}

	found := false

	for idx := range todos {
		if todos[idx].ID == id {
			todos[idx].Title = title
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

	fmt.Printf("Todo %d edited successfully!\n", id)
}
