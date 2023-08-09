package af

import (
	"fmt"
	"os"
)

func ExecuteCreateFile(args []string) {
	if len(args) != 1 {
		fmt.Println("Invalid arguments for 'af'. Use the help command to see proper use.")
		return
	}

	filename := args[0]
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	fmt.Println("File created:", filename)
}
