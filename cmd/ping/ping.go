package ping

import (
	"fmt"
	"os/exec"
	"runtime"
)

// ! Temperary fix and should be made custom in the future
func ExecutePing(args []string) {
	if len(args) != 2 {
		fmt.Println("Usage: ping <host>")
		return
	}

	host := args[1]
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("ping", host)
	} else {
		cmd = exec.Command("ping", "-c", "4", host) // For Linux and macOS
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running ping command:", err)
		return
	}
	fmt.Println(string(output))
}
