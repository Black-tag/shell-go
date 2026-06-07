package shell

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/codecrafters-io/shell-starter-go/app/internal/builtins"
	"github.com/codecrafters-io/shell-starter-go/app/internal/parser"
)

func (s *Shell) Execute(cmd parser.Command) {

	switch cmd.Name {

	case "exit":

		os.Exit(0)

	case "echo":
		builtins.Echo(cmd.Args)

	case "cd":
		builtins.Cd(cmd.Args)

	case "pwd":
		pwd, err := builtins.Pwd()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(pwd)
		
	case "type":
		builtins.Type(cmd.Args)

	default:
		// run external command
		path, err := exec.LookPath(cmd.Name)

		if err != nil {
			fmt.Printf("%s: command not found\n", cmd.Name)
			return

		}
		command := exec.Command(path, cmd.Args...)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		command.Stdin = os.Stdin

		command.Run()
	}

}
