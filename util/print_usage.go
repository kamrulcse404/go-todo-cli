package util

import (
	"fmt"
)

func PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println("  todo add \"Task\"")
	fmt.Println("  todo list")
	fmt.Println("  todo complete <id>")
	fmt.Println("  todo delete <id>")
	fmt.Println("  todo edit <id> \"Title\"")
	fmt.Println("  todo search <id>")
}