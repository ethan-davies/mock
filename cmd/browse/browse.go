package browse

import (
	"fmt"

	"github.com/skratchdot/open-golang/open"
)

func ExecuteBrowse(args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: browser <URL>")
		return
	}

	url := args[0]
	err := open.Run(url)
	if err != nil {
		fmt.Println("Error opening browser:", err)
	}
}
