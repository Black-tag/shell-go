package parser


import (
	"strings"


)


func Parse(input string) Command {
	var tokens []string
	var current strings.Builder

	inSingleQuotes := false
	inDoubleQuotes := false

	for _, ch := range input {

	
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

			default:
				current.WriteRune(ch)
			}
	}

	if current.Len() > 0 {
		tokens = append(tokens, current.String())
	}

	return Command{
		Name: tokens[0],
		Args: tokens[1:],
	}

}
