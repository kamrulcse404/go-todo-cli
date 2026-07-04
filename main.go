package main

import (
	"fmt"
	"os"
	"todo-cli/commands"
	"todo-cli/util"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("todo add \"Task\"")
		fmt.Println("todo list")
		fmt.Println("todo complete <id>")
		fmt.Println("todo delete <id>")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a todo title.")
			return
		}
		title := os.Args[2]
		commands.AddTodo(title)
	case "list":
		commands.ListTodos()
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a todo ID.")
			return
		}

		id, err := util.StringToInt(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID.")
			return
		}
		commands.CompleteTodo(id)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a todo ID.")
			return
		}
		id, err := util.StringToInt(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID.")
			return
		}
		commands.DeleteTodo(id)
	}
}
