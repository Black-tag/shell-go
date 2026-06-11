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
	
	

	if cmd.IsBackgorund {

		command := exec.Command(cmd.Name, cmd.Args...)

		command.Stdout = stdout
		command.Stderr = stderr
		command.Stdin = os.Stdin

		err := command.Start()
		if err != nil {
			fmt.Println(err)
			return
		}

		job := &Job{
			ID: s.NextJobID,
			PID: command.Process.Pid,
			Command: cmd.Name,
			Done: false,
		}
		s.Jobs = append(s.Jobs, job)
		s.NextJobID++
		fmt.Printf("[%d] %d\n", job.ID, job.PID)

		go func(j *Job, c *exec.Cmd) {
			err := c.Wait()
			if err != nil {
				j.Done = true
			}
		}(job, command)

		return
		
	}
	
	

	if cmd.StdoutRedirect != "" {
		var file *os.File
		var err error

		if cmd.StdoutAppend {
			file, err = os.OpenFile(
				cmd.StdoutRedirect,
				os.O_CREATE|os.O_WRONLY|os.O_APPEND,
				0644,
			)
		} else {
			file, err = os.Create(cmd.StdoutRedirect)
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()
		stdout = file
	}
	
	if cmd.StderrRedirect != "" {
		var file *os.File
		var err error

		if cmd.StderrAppend {
			file, err = os.OpenFile(
				cmd.StderrRedirect,
				os.O_CREATE|os.O_WRONLY|os.O_APPEND,
				0644,
			)
		} else {
			file, err = os.Create(cmd.StderrRedirect)
		}

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
