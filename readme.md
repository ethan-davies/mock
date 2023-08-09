# Mock Shell

Mock Shell is a simple command-line shell program implemented in Go. It provides basic shell commands and functionalities to interact with the file system and execute system commands.

## Features

- Change current directory (`cd`)
- List files in the current directory (`ls` or `dir`)
- Clear the terminal screen (`clear` or `cls`)
- Display text (`echo`)
- Print current directory (`pwd`)
- Display current user (`whoami`)
- Create a new file (`af <filename>`)
- Create a new directory (`mkdir <dirname>`)
- Copy a file (`cp <source> <destination>`)
- Move or rename a file/directory (`mv <source> <destination>`)
- Remove a file or directory (`rm <path>`)
- Remove an empty directory (`rmdir <dirname>`)
- Display current date and time (`time`)
- Display disk usage for the current directory (`du` or `diskusage`)
- Display process status (`ps`)
- Exit the shell (`exit`)

## Getting Started

1. Clone this repository:

   ```bash
   git clone https://github.com/ethan-davies/mock.git
   cd mock
   ```

## Building
For build to different platform use these: 

go build ./internal/shell

**Microsoft Powershell**
```bash
$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o bin/mock-shell-windows.exe shell/main.go
$env:GOOS="darwin"; $env:GOARCH="amd64"; go build -o bin/mock-shell-darwin shell/main.go
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o bin/mock-shell-linux shell/main.go
```