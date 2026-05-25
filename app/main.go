package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)





var _ = fmt.Print
func isBuiltin(command string, wholeCommand []string) string {
	builtIns := []string{"echo", "exit", "type"}
	for _, builtIn := range builtIns {
		if command  == builtIn {
			
			return fmt.Sprintf("%s is a shell builtin", command)
		}
		
	}

	path , err := exec.LookPath(command)
	if err == nil {
		return fmt.Sprintf("%s is %s", command, path)
	}
	return fmt.Sprintf("%s: not found", command)
	
	}
func getPwd(command string) string {

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return pwd
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
		case "pwd":
			getPwd(command)


		default:
			_, err := exec.LookPath(commandArray[0])
			if err != nil {
				fmt.Println(command + ": command not found")
			} else {
				cmd := exec.Command(commandArray[0], commandArray[1:]...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Stdin = os.Stdin
				cmd.Run()
			}

		}
		


	}
	
	
}

