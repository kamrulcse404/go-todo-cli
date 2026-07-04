package commands

import (
	"todo-cli/models"
	"todo-cli/storage"
	"fmt"
)

func AddTodo(title string) {
	todos, err := storage.LoadTodos()
	if err != nil {
		fmt.Println("Error loading todos!", err)
		return 
	}

	newTodo := models.Todo{
		ID:         GetNextID(todos),
		Title:     title,
		Completed: false,
	}

	todos = append(todos, newTodo)

	err = storage.SaveTodos(todos)
	if err != nil {
		fmt.Println("Error adding todo! ", err)
		return
	}

	fmt.Println("Todo added successfully!")
}
