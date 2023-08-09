package rm

import (
	"fmt"
	"os"
)

func ExecuteRm(args []string) {
	if len(args) != 1 {
		fmt.Println("Invalid usage of 'rm'. Use the help command to see proper use.")
		return
	}

	path := args[0]
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Println("Error removing:", err)
		return
	}

	fmt.Println("Removed:", path)
}
