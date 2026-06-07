package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/internal/parser"
	"github.com/codecrafters-io/shell-starter-go/app/internal/shell"
	// "github.com/codecrafters-io/shell-starter-go/app/internal/builtins"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	sh := shell.New()

	for {
		fmt.Print("$ ")

		input, err := reader.ReadString('\n')
		if err != nil {
			os.Exit(1)
		}

		cmd := parser.Parse(strings.TrimSpace(input))

		if cmd.Name == "" {
			continue
		}
		sh.Execute(cmd)
	}

}
