package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var _ = fmt.Print

func parseCommand(input string) []string {
	var tokens []string
	var current strings.Builder

	inSingleQuotes := false

	for _, ch := range input {

		switch {
		case ch == '\'':
			inSingleQuotes = !inSingleQuotes

		case (ch == ' ' || ch == '\t') && !inSingleQuotes:
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}

		default:
			current.WriteRune(ch)

		}
	}

	if current.Len() > 0 {
		tokens = append(tokens, current.String())
	}

	return tokens

}

func isBuiltin(command string, wholeCommand []string) string {
	builtIns := []string{"echo", "exit", "type", "pwd", "cd"}
	for _, builtIn := range builtIns {
		if command == builtIn {

			return fmt.Sprintf("%s is a shell builtin", command)
		}

	}

	path, err := exec.LookPath(command)
	if err == nil {
		return fmt.Sprintf("%s is %s", command, path)
	}
	return fmt.Sprintf("%s: not found", command)

}
func getPwd() string {

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	return pwd
}
func dealCd(commandArray []string) {
	if len(commandArray) < 2 {
		return
	}

	dirToChange := commandArray[1]
	if dirToChange == "~" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Could not find home directory for user")
			return
		}

		dirToChange = homeDir
	}

	info, err := os.Stat(dirToChange)

	if err != nil || !info.IsDir() {
		fmt.Printf("cd: %s: No such file or directory\n", dirToChange)
		return
	}

	err = os.Chdir(dirToChange)
	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", dirToChange)
	}
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

		commandArray := parseCommand(command)
		if len(commandArray) == 0 {
			continue
		}
		switch commandArray[0] {
		case "exit":
			return

		case "echo":
			fmt.Println(strings.Join(commandArray[1:], " "))
		case "type":
			fmt.Println(isBuiltin(commandArray[1], commandArray))
		case "pwd":
			fmt.Println(getPwd())
		case "cd":
			dealCd(commandArray)

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
