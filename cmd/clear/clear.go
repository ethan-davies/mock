package clear

import (
	"os"
	"os/exec"
	"runtime"
)

func ExecuteClear() {
	clearCmd := exec.Command("clear") // Unix-like systems
	if runtime.GOOS == "windows" {
		clearCmd = exec.Command("cmd", "/c", "cls") // Windows
	}

	clearCmd.Stdout = os.Stdout
	clearCmd.Run()
}
