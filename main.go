package main

import (
	"fmt"
	"os"
	"strings"
	"todo-cli/commands"
	"todo-cli/util"
)

func main() {
	if len(os.Args) < 2 {
		util.PrintUsage()
		return
	}
	command := strings.ToLower(os.Args[1])

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a todo title.")
			util.PrintUsage()
			return
		}
		title := os.Args[2]
		commands.AddTodo(title)

	case "list":
		commands.ListTodos()

	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a todo ID.")
			util.PrintUsage()
			return
		}

		id, err := util.StringToInt(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID. Please provide a numeric ID.")
			return
		}
		commands.CompleteTodo(id)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a todo ID.")
			util.PrintUsage()
			return
		}
		id, err := util.StringToInt(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID.")
			util.PrintUsage()
			return
		}
		commands.DeleteTodo(id)

	case "edit":
		if len(os.Args) < 4 {
			fmt.Println("Usage: todo edit <id> \"Title\"")
			util.PrintUsage()
			return
		}
		title := os.Args[3]
		id, err := util.StringToInt(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID.")
			util.PrintUsage()
			return
		}
		commands.EditTodo(id, title)

	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a todo ID.")
			util.PrintUsage()
			return
		}
		id, err := util.StringToInt(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID.")
			util.PrintUsage()
			return
		}
		todo, err := commands.SearchTodo(id)
		if err != nil {
			fmt.Println(err)
			return
		}
		status := "✗"
		if todo.Completed {
			status = "✓"
		}
		fmt.Printf("%d. [%s] %s\n", todo.ID, status, todo.Title)

	case "pending":
		commands.PeningTodosList()

	case "completed":
		commands.CompletedTodosList()

	default:
		fmt.Println("Unknown command.")
		util.PrintUsage()
	}
}
