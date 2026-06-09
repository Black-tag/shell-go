package parser

import (
	"strings"

	
)

func Parse(input string) Command {
	var tokens []string
	var current strings.Builder
	var redirect string

	inSingleQuotes := false
	inDoubleQuotes := false

	for i := 0; i< len(input); i++ {
		if tokens[i] == ">" {
			if i+1 < len(tokens) {
			redirect = tokens[i+1]
			}
			i++
			continue

		}
		
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
	if len(tokens) == 0 {
		return Command{}
	}

	return Command{
		Name: tokens[0],
		Args: tokens[1:],
		stderrRedirect: redirect,
	}

}
