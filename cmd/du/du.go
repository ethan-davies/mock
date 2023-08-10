package du

import (
	"fmt"
	"os"
	"path/filepath"
)

func ExecuteDU(args []string) {
	var dir string

	if len(args) == 0 {
		// No arguments provided, use the current directory
		dir, _ = os.Getwd()
	} else {
		dir = args[0]
	}

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
