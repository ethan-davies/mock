package rmdir

import (
	"fmt"
	"os"
)

func ExecuteRmdir(args []string) {
	if len(args) != 1 {
		fmt.Println("Invalid usage of 'rmdir'. Use the help command to see proper use.")
		return
	}

	dir := args[0]
	err := os.Remove(dir)
	if err != nil {
		fmt.Println("Error removing directory:", err)
		return
	}

	fmt.Println("Directory removed:", dir)
}
