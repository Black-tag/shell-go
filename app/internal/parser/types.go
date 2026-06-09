package parser

type Command struct {
	Name string
	Args []string
	StdoutRedirect string
	stderrRedirect string
}
