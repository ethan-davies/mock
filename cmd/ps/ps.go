package ps

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecuteProcessStatus() {
	tasklistCmd := exec.Command("tasklist")
	tasklistCmd.Stdout = os.Stdout
	tasklistCmd.Stderr = os.Stderr
	err := tasklistCmd.Run()
	if err != nil {
		fmt.Println("Error displaying process status:", err)
	}
}
