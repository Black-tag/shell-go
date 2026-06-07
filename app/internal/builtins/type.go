package builtins

import (
	"fmt"
	"os/exec"
)



var Builtins = map[string]bool{
	"echo": true,
	"exit": true,
	"type": true,
	"pwd": true,
	"cd": true,

}

func Type(args []string) {
	target := args[0]

	if Builtins[target] {
		fmt.Printf("%s is a shell builtin\n", target)
		return
	}

	path, err := exec.LookPath(target)
	if err == nil {
		fmt.Printf("%s is %s\n", target, path)
		return
	}

	fmt.Printf("%s: not found\n", target)
}