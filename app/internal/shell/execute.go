package shell

import (
	"github.com/codecrafters-io/shell-starter-go/app/internal/builtins"
	"github.com/codecrafters-io/shell-starter-go/app/internal/parser"
)





func (s *Shell) Execute(cmd parser.Command) {
	

	switch cmd.Name {
		

	case "exit":
		// ned to fill

	case "echo":
		// later will fill
		builtins.Echo(cmd.Args)

	default:
		// nohting for now 
	}

}