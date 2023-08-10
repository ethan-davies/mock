package whoami

import (
	"fmt"
	"os/user"
)

func ExecuteWhoAmI() {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Username:", usr.Username)
}
