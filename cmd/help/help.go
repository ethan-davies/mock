package help

import "fmt"

func ExecuteHelp() {
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
	fmt.Println("browser          Open a URL in the users prefered browser")
}