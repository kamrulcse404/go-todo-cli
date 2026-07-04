package commands

import (
	"fmt"
	"todo-cli/models"
	"todo-cli/storage"
)

func AddTodo(title string) {
	todos, err := storage.LoadTodos()
	if err != nil {
		fmt.Println("Error loading todos!", err)
		return
	}

	if title == "" {
		fmt.Println("Todo title cannot be empty.")
		return
	}

	newTodo := models.Todo{
		ID:        GetNextID(todos),
		Title:     title,
		Completed: false,
	}

	todos = append(todos, newTodo)

	err = storage.SaveTodos(todos)
	if err != nil {
		fmt.Println("Error adding todo! ", err)
		return
	}
	
	fmt.Printf("Todo #%d added successfully.\n", newTodo.ID)
}
