package cp

import (
	"fmt"
	"io"
	"os"
)

func ExecuteCP(srcPath, destPath string) {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer srcFile.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		fmt.Println("Error copying file:", err)
		return
	}

	fmt.Println("File copied successfully")
}
