package builtins

import "fmt"



var BUiltins = map[string]bool{
	"echo": true,
	"exit": true,
	"type": true,
	"pwd": true,
	"cd": true,

}

func Type(args []string) {
	if BUiltins[args[0]] {
		fmt.Printf("%s is a shell builtin\n", args[0])
    	return

	}

}