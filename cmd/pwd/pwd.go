package pwd

import (
	"fmt"
	"os"
)

func ExecutePwd(args []string) {
	if len(args) != 1 {
		fmt.Println("Invalid usage of 'pwd'. Use the help command to see proper use.")
		return
	}

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(currentDir)
}
