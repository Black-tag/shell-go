package parser

import (
	
	"strings"
)

func Parse(input string) Command {
	var tokens []string
	var current strings.Builder
	

	inSingleQuotes := false
	inDoubleQuotes := false

	for i := 0; i< len(input); i++ {
		
		ch := rune(input[i])

		switch {

		case ch == '\'' && !inDoubleQuotes:
			inSingleQuotes = !inSingleQuotes

		case ch == '"' && !inSingleQuotes:
			inDoubleQuotes = !inDoubleQuotes

		case (ch == ' ' || ch == '\t') &&
			!inSingleQuotes &&
			!inDoubleQuotes:

			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}
		case ch == '\\' && !inSingleQuotes && !inDoubleQuotes:
			if i+1 < len(input) {
				next := rune(input[i+1])
				current.WriteRune(next)
				i++
			}
		case ch == '\\' && inDoubleQuotes:
			if i+1 < len(input) {
				next := rune(input[i+1])

				if next == '\\' || next == '"' {
					current.WriteRune(next)
					i++
				} else {
					current.WriteRune('\\')
				}
			}
		
			

		default:
			current.WriteRune(ch)
		}
	}

	if current.Len() > 0 {
		tokens = append(tokens, current.String())
	}
	var args []string

	var stdoutRedirect string
	var stderrRedirect string

	for i := 0; i < len(tokens); i++ {

		switch tokens[i] {

		case ">", "1>":
			if i+1 < len(tokens) {
				stdoutRedirect = tokens[i+1]
			}
			i++

		case "2>":
			if i+1 < len(tokens) {
				stderrRedirect = tokens[i+1]
			}
			i++

		default:
			args = append(args, tokens[i])
		}
	}

	if len(args) == 0 {
		return Command{}
	}
	
	

	return Command{
		Name: args[0],
		Args: args[1:],
		StdoutRedirect: stdoutRedirect,
		StderrRedirect: stderrRedirect,
	}

}
