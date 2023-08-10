package mv

import (
	"fmt"
	"os"
)

func ExecuteMV(args []string) {
	if len(args) != 3 {
		fmt.Println("Usage: mv <source> <destination>")
		return
	}

	source := args[1]
	destination := args[2]

	err := os.Rename(source, destination)
	if err != nil {
		fmt.Println("Error moving file:", err)
	} else {
		fmt.Println("File moved successfully")
	}
}
