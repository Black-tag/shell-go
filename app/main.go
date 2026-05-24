package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)




type Cmd struct {
	path string
	args []string
}
var _ = fmt.Print
func isBuiltin(command string, wholeCommand []string) string {
	builtIns := []string{"echo", "exit", "type"}
	for _, builtIn := range builtIns {
		if command  == builtIn {
			
			return fmt.Sprintf("%s is a shell builtin", command)
		}
		
	}

	path , err := exec.LookPath(command)
	

	cmd := Cmd{
		path: path,
		args: wholeCommand,

	}
	
	if err == nil {
		cmd := exec.Command(cmd.path, cmd.args...)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		err = cmd.Run()

		if err != nil {
			fmt.Println(err)
		}

	}
	return fmt.Sprintf("%s: not found", command)
	
	}

func main() {

	
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")

	
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		command = strings.TrimSpace(command)
		
		commandArray := strings.Split(command, " ")
		switch commandArray[0] {
		case "exit":
			return
		
		case "echo":
			fmt.Println(strings.Join(commandArray[1:], " "))
		case "type":
			fmt.Println(isBuiltin(commandArray[1], commandArray))

		default:
			fmt.Println(command + ": command not found")

		}
		


	}
	
	
}

