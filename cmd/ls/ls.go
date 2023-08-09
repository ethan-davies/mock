package ls

import (
	"fmt"
	"os"
)

func ExecuteLS(args []string) {
	dir := "."
	if len(args) == 2 {
		dir = args[1]
	} else if len(args) > 2 {
		fmt.Println("Invalid arguments for 'ls'. Use the help command to see proper use.")
		return
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
