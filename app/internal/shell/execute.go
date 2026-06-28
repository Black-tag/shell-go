package shell

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/internal/builtins"
	"github.com/codecrafters-io/shell-starter-go/app/internal/parser"
)

func (s *Shell) Execute(
	cmd parser.Command,
	cstdin io.Reader,
	cstdout io.Writer,
	cstderr io.Writer) {
	var stdin io.Reader = cstdin
	var stdout io.Writer = cstdout
	var stderr io.Writer = cstderr

	if cmd.IsBackgorund {

		command := exec.Command(cmd.Name, cmd.Args...)

		command.Stdout = stdout
		command.Stderr = stderr
		command.Stdin = stdin

		err := command.Start()
		if err != nil {
			fmt.Println(err)
			return 
		}

		job := &Job{
			ID:  s.getNextJobID(),
			PID: command.Process.Pid,
			Command: strings.Join(
				append([]string{cmd.Name}, cmd.Args...),
				" ",
			) + " &",
			Status: "Running",
			Cmd:    *command,
		}
		s.mu.Lock()
		s.Jobs = append(s.Jobs, job)
		if len(s.Jobs) == 0 {
			s.NextJobID = 1
		}
		
		s.NextJobID++
		s.mu.Unlock()
		fmt.Printf("[%d] %d\n", job.ID, job.PID)

		go func(j *Job, c *exec.Cmd) {
			_ = c.Wait()
			s.mu.Lock()
			j.Command = strings.TrimSuffix(j.Command, " &")
			j.Status = "Done"
			s.mu.Unlock()
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
		fmt.Fprintln(stdout, pwd)
	// case "jobs":
	// 	return

	case "type":
		builtins.Type(cmd.Args , stdout)

	case "jobs":
		// fmt.Print("switching case inside jobs")
		s.Job()

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
		command.Stdin = stdin

		command.Run()
	}

}


func (s *Shell)ExecutePipeline(pipeline parser.Pipeline) {

	// left := pipeline.Commands[0]
	// right := pipeline.Commands[1]
	previousReader := io.Reader(os.Stdin)

	for i, cmd := range pipeline.Commands {
	
		if i == len(pipeline.Commands)-1 {
			s.Execute(cmd, previousReader, os.Stdout, os.Stderr)
			break
		}else {
			r, w := io.Pipe()

			
			go func (
				cmd parser.Command,
				stdin io.Reader,
				stdout *io.PipeWriter,
			)  {
				defer stdout.Close()
				s.Execute(cmd, stdin, stdout, os.Stderr)
				w.Close()
				
			}(cmd, previousReader, w)
			previousReader = r
		}

	}


	
	

}
