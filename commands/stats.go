package commands

import(
	"todo-cli/storage"
	"fmt"
)

func TodoStats() {
	todos, err := storage.LoadTodos()

	if err != nil {
		fmt.Println("Error loading todos:", err)
		return
	}

	total := len(todos)
	completed := 0

	for _, todo := range todos {
		if todo.Completed  {
			completed++
		}
	}

	pending := total - completed
	var percent float64

	if total > 0 {
		percent = (float64(completed) / float64(total)) * 100
	}

	fmt.Println("Todo Stats")
	fmt.Println("----------------")
	fmt.Println("Total     :", total)
	fmt.Println("Completed :", completed)
	fmt.Println("Pending   :", pending)
	fmt.Printf("Progress  : %.2f%%\n", percent)
}