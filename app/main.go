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
		sh.ReapJobs()
		fmt.Print("$ ")

		input, err := reader.ReadString('\n')
		if err != nil {
			os.Exit(1)
		}

		pipeline := parser.ParsePipeline(strings.TrimSpace(input))

		if len(pipeline.Commands) == 1 {
			sh.Execute(pipeline.Commands[0], os.Stdin, os.Stdout, os.Stderr)
			if pipeline.Commands[0].Name == "" {
				continue
			}
			sh.ReapJobs()
		}else {
			sh.ExecutePipeline(pipeline)
		}

		// if cmd.Name == "" {
		// 	continue
		// }
		// sh.Execute(cmd)
		// sh.ReapJobs()
		// sh.ExecutePipeline(pipeline)
	}

}
