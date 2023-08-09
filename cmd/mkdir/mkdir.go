package mkdir

import (
	"fmt"
	"os"
)

func ExecuteMkdir(args []string) {
	if len(args) != 2 {
		fmt.Println("Invalid arguments for 'mkdir'. Use the help command to see proper use.")
		return
	}

	dirname := args[1]
	err := os.Mkdir(dirname, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	fmt.Println("Directory created:", dirname)
}
