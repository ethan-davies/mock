package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)


func addToPathForShell(executables ...string) {
	for _, exe := range executables {
		path := os.Getenv("PATH")
		path = path + string(filepath.ListSeparator) + exe
		os.Setenv("PATH", path)
	}
}

var validCommands = []string{
	"cd", "ls", "dir", "clear", "cls", "help", "pwd",
	"echo", "whoami", "af", "mkdir", "time", "du", "diskusage",
	"ps", "exit", "rm", "rmdir", "ping", "cp", "mv",
}

func main() {

	// Add paths to executables (if needed)
	addToPathForShell("C:\\path\\to\\git\\bin", "C:\\path\\to\\docker", "/usr/local/bin") // Example paths

	for {
        // Get username and hostname
        user, _ := user.Current()
        hostname, _ := os.Hostname()
		currentDir, _ := os.Getwd()

        // Format the prompt
        prompt := fmt.Sprintf("%s@%s:%s $ ", user.Username, hostname, currentDir)

        // Print the prompt
        fmt.Print(prompt)

        // Read input and process commands
        reader := bufio.NewReader(os.Stdin)
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

	switch {
	case strings.HasPrefix(input, "cd "):
		if len(strings.TrimSpace(input[3:])) == 0 {
			fmt.Println("Invalid arguments. Use the help command to see proper use.")
		} else {
			changeDirectory(strings.TrimSpace(input[3:]))
		}
	case input == "ls" || input == "dir":
		listFiles()
	case input == "clear" || input == "cls":
		clearScreen()
	case input == "help":
		showHelp()
	case input == "pwd":
		fmt.Println(currentDir)

	case strings.HasPrefix(input, "echo "):
		args := strings.TrimSpace(input[5:])
		if len(args) == 0 {
			fmt.Println("Invalid arguments. Use the help command to see proper use.")
		} else {
			echoText(args)
		}

	case input == "whoami":
		whoAmI()
	case strings.HasPrefix(input, "af "):
		args := strings.Fields(input[3:])
		if len(args) != 1 {
			fmt.Println("Invalid arguments. Use the help command to see proper use.")
		} else {
			createFile(args[0])
		}
	case strings.HasPrefix(input, "mkdir "):
		if len(strings.TrimSpace(input[6:])) == 0 {
			fmt.Println("Invalid arguments. Use the help command to see proper use.")
		} else {
			createDirectory(strings.TrimSpace(input[6:]))
		}
	case input == "time":
		showTime()
	case input == "du" || input == "diskusage":
		showDiskUsage(currentDir)
	case input == "ps":
		showProcesses()
	case strings.HasPrefix(input, "rm "):
		if len(strings.TrimSpace(input[3:])) == 0 {
			fmt.Println("Invalid arguments. Use the help command to see proper use.")
		} else {
			args := strings.TrimSpace(input[3:])
			err := removeFileOrDirectory(args)
			if err != nil {
				fmt.Println("Error removing:", err)
			}
		}
	case strings.HasPrefix(input, "rmdir "):
		if len(strings.TrimSpace(input[6:])) == 0 {
			fmt.Println("Invalid arguments. Use the help command to see proper use.")
		} else {
			dir := strings.TrimSpace(input[6:])
			err := removeDirectory(dir)
			if err != nil {
				fmt.Println("Error removing directory:", err)
			} else {
				fmt.Println("Directory removed:", dir)
			}
		}
		
	case strings.HasPrefix(input, "ping "):
		host := strings.TrimSpace(input[5:])
		pingHost(host)

	case strings.HasPrefix(input, "cp "):
		args := strings.Fields(strings.TrimSpace(input[3:]))
		if len(args) != 2 {
			fmt.Println("Invalid arguments. Use the help command to see proper use.")
		} else {
			err := copyFile(args[0], args[1])
			if err != nil {
				fmt.Println("Error copying file:", err)
			} else {
				fmt.Println("File copied successfully")
			}
		}




	case strings.HasPrefix(input, "mv "):
		args := strings.SplitN(strings.TrimSpace(input[3:]), " ", 2)
		if len(args) != 2 {
			fmt.Println("Usage: mv <source> <destination>")
		} else {
			err := moveFile(args[0], args[1])
			if err != nil {
				fmt.Println("Error moving file:", err)
			} else {
				fmt.Println("File moved successfully")
			}
		}
	default:
		runCommand(input)
	}

	}
}


func runCommand(input string) {
	if !strings.ContainsAny(input, " \t") {
		cmd := exec.Command(input)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			if contains(validCommands, input) {
				fmt.Println("Invalid arguments. Use the help command to see proper use.")
			} else {
				fmt.Println("Command not found:", input)
			}
		}
	} else {
		fmt.Println("Command not found:", input)
	}
}

func contains(list []string, target string) bool {
    for _, item := range list {
        if item == target {
            return true
        }
    }
    return false
}

func changeDirectory(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func listFiles() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func clearScreen() {
	clearCmd := exec.Command("clear") // Unix-like systems
	if runtime.GOOS == "windows" {
		clearCmd = exec.Command("cmd", "/c", "cls") // Windows
	}

	clearCmd.Stdout = os.Stdout
	clearCmd.Run()
}

func showHelp() {
	fmt.Println("Available commands:")
	fmt.Println("cd <directory>   Change current directory")
	fmt.Println("ls or dir        List files in the current directory")
	fmt.Println("clear or cls     Clear the terminal screen")
	fmt.Println("echo <text>      Display text")
	fmt.Println("pwd              Print current directory")
	fmt.Println("whoami           Display current user")
	fmt.Println("af <filename>    Create a new file")
	fmt.Println("mkdir <dirname>  Create a new directory")
	fmt.Println("cp <source> <destination>  Copy a file")
	fmt.Println("mv <source> <destination>  Move or rename a file/directory")
	fmt.Println("rm <path>        Remove a file or directory")
	fmt.Println("rmdir <dirname>  Remove an empty directory")
	fmt.Println("time             Display current date and time")
	fmt.Println("du or diskusage  Display disk usage for the current directory")
	fmt.Println("ps               Display process status")
	fmt.Println("exit             Exit the shell")
}

func echoText(text string) {
	fmt.Println(text)
}

func whoAmI() {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Username:", usr.Username)
}

func createFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	fmt.Println("File created:", filename)
}

func createDirectory(dirname string) {
	err := os.Mkdir(dirname, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	fmt.Println("Directory created:", dirname)
}

func showTime() {
	currentTime := time.Now()
	fmt.Println("Current Date and Time:", currentTime.Format("2006-01-02 15:04:05"))
}

func showDiskUsage(dir string) {
	totalSize := calculateDiskUsage(dir)
	fmt.Printf("Disk Usage for %s: %s\n", dir, formatBytes(totalSize))
}

func calculateDiskUsage(dir string) int64 {
	var totalSize int64
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				fmt.Println("Error:", err)
				return nil
			}
			totalSize += info.Size()
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error calculating disk usage:", err)
	}
	return totalSize
}

func formatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func showProcesses() {
	tasklistCmd := exec.Command("tasklist")
	tasklistCmd.Stdout = os.Stdout
	tasklistCmd.Stderr = os.Stderr
	err := tasklistCmd.Run()
	if err != nil {
		fmt.Println("Error displaying process status:", err)
	}
}

func copyFile(srcPath, destPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

func moveFile(srcPath, destPath string) error {
	err := os.Rename(srcPath, destPath)
	if err != nil {
		return err
	}

	return nil
}

func pingHost(host string) {
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

func removeFileOrDirectory(path string) error {
	err := os.RemoveAll(path)
	return err
}

func removeDirectory(dir string) error {
	err := os.RemoveAll(dir)
	return err
}


// Make each command a different file in ./mock/shell/commands
