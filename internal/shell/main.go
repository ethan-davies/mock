package main

import (
	"bufio"
	"fmt"
	"mock-shell/cmd/af"
	"mock-shell/cmd/browse"
	"mock-shell/cmd/cd"
	"mock-shell/cmd/clear"
	"mock-shell/cmd/cp"
	"mock-shell/cmd/du"
	"mock-shell/cmd/echo"
	"mock-shell/cmd/exit"
	"mock-shell/cmd/help"
	"mock-shell/cmd/ls"
	"mock-shell/cmd/mkdir"
	"mock-shell/cmd/mv"
	"mock-shell/cmd/ping"
	"mock-shell/cmd/ps"
	"mock-shell/cmd/pwd"
	"mock-shell/cmd/rm"
	"mock-shell/cmd/rmdir"
	"mock-shell/cmd/time"
	"mock-shell/cmd/whoami"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func getUsername() string {
	user, _ := user.Current()
	parts := strings.Split(user.Username, "\\")
	return parts[1] + "@" + parts[0]
}


func main() {
	for {
		currentDir, _ := os.Getwd()

		// Get the path relative to the root directory
		rootDir, _ := filepath.Abs("/")
		relPath, _ := filepath.Rel(rootDir, currentDir)
		relPath = strings.Replace(relPath, "\\", "/", -1)
		
		// Construct the prompt directly using getUsername()
		prompt := getUsername() + ":" + relPath + " $ "
		
		fmt.Print(prompt)

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		// Process the command
		processCommand(input)
	}
}

func processCommand(input string) {
	args := strings.Fields(input) // Split input into command and arguments
	command := args[0]
	remainingArgs := args[1:]

	if len(args) == 0 {
		return
	}

	switch command {
	case "cd":
		if len(args) == 2 {
			cd.ExecuteCD(args[1])
		} else {
			fmt.Println("Invalid arguments for 'cd'. Instead use 'cd <destination>' or use the help command to see proper use.")
		}
	case "clear", "cls":
		clear.ExecuteClear()
	case "cp":
		if len(args) != 3 {
			fmt.Println("Invalid arguments for 'cp'. Instead use 'cp <source> <destination>' or use the help command to see proper use.")
		} else {
			cp.ExecuteCP(args[1], args[2])
		}
	case "echo":
		if len(args) >= 2 {
			text := strings.Join(args[1:], " ")
			echo.ExecuteEcho(text)
		} else {
			fmt.Println("Invalid arguments for 'echo'. Instead use 'echo <text>' or use the help command to see proper use.")
		}
	case "exit":
		exit.ExecuteExit()
	case "ls", "dir":
		ls.ExecuteLS(args)
	case "mkdir":
		mkdir.ExecuteMkdir(args)
	case "mv":
		if len(args) != 3 {
			fmt.Println("Invalid arguments for 'mv'. Instead use 'mv <source> <destination>' or use the help command to see proper use.")
		} else {
			mv.ExecuteMV(args)
		}
	case "ping":
		if len(args) != 2 {
			fmt.Println("Invalid arguments for 'ping'. Instead use 'ping <host>' or use the help command to see proper use.")
		} else {
			ping.ExecutePing(args)
		}
	case "pwd": 
		pwd.ExecutePwd(args)
	case "rm":
		rm.ExecuteRm(args[1:])
	case "rmdir":
		rmdir.ExecuteRmdir(args[1:])
	case "time":
		time.ExecuteTime()
	case "du", "diskusage":
		du.ExecuteDU(remainingArgs)
	case "ps":
		ps.ExecuteProcessStatus()
	case "af":
		af.ExecuteCreateFile(args[1:])
	case "whoami":
		whoami.ExecuteWhoAmI()
	case "help":
		help.ExecuteHelp()
	case "browse":
		browse.ExecuteBrowse(args[1:])
	default:
		fmt.Printf("Command not found: %s. Use the help command to see proper use.\n", command)
	}
}
