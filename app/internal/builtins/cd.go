package builtins

import (
	"fmt"
	"os"
)

func Cd(args []string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Could not determine home directory")
		return
	}

	var dirToChange string

	switch len(args) {
	case 0:
		dirToChange = homeDir

	case 1:
		if args[0] == "~" {
			dirToChange = homeDir
		} else {
			dirToChange = args[0]
		}

	default:
		fmt.Println("cd: too many arguments")
		return
	}

	err = os.Chdir(dirToChange)
	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", dirToChange)
		return
	}
}
