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
