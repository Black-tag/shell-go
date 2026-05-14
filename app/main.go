package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	// TODO: Uncomment the code below to pass the first stage
	
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")

		// read - step 1
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

		default:
			fmt.Println(command + ": command not found")

		}
		


	}
	
}
