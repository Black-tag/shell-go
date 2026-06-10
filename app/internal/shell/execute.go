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
	var stdout io.Writer = os.Stdout
	var stderr io.Writer = os.Stderr
	
	
	

	if cmd.StdoutRedirect != "" {
		file, err := os.Create(cmd.StdoutRedirect)
		if err != nil {
			fmt.Println(err)
			return
		}
		
		defer file.Close()
		stdout = file
		
	}
	if cmd.StderrRedirect  != ""{
		file, err := os.Create(cmd.StderrRedirect)
		if err != nil {
			fmt.Println(err)
			return
		}
		
		defer file.Close()
		stderr = file
		
	}



		

	switch cmd.Name {

	case "exit":

		os.Exit(0)

	case "echo":
		builtins.Echo(cmd.Args, stdout)

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

		
		
		command.Stdout = stdout
		command.Stderr = stderr
		command.Stdin = os.Stdin

		command.Run()
	}

}
