package time

import (
	"fmt"
	"time"
)

func ExecuteTime() {
	currentTime := time.Now()
	fmt.Println("Current Date and Time:", currentTime.Format("2006-01-02 15:04:05"))
}
