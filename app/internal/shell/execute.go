package shell

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/codecrafters-io/shell-starter-go/app/internal/builtins"
	"github.com/codecrafters-io/shell-starter-go/app/internal/parser"
)

func (s *Shell) Execute(cmd parser.Command) {
	var out io.Writer = os.Stdout
	// var output *os.File

		if cmd.StdoutRedirect != "" {
			file, err := os.Create(cmd.StdoutRedirect)
			if err != nil {
				fmt.Println(err)
				return
			}
			
			defer file.Close()
			out = file
			
		} 

		

	switch cmd.Name {

	case "exit":

		os.Exit(0)

	case "echo":
		builtins.Echo(cmd.Args, out)

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
		_, err := exec.LookPath(cmd.Name)

		if err != nil {
			fmt.Printf("%s: command not found\n", cmd.Name)
			return

		}
		command := exec.Command(cmd.Name, cmd.Args...)

		
		
		command.Stdout = out
		command.Stderr = os.Stderr
		command.Stdin = os.Stdin

		command.Run()
	}

}
