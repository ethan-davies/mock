package cd

import (
	"fmt"
	"os"
)

func ExecuteCD(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
