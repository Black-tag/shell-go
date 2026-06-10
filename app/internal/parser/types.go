package parser

type Command struct {
	Name string
	Args []string
	StdoutRedirect string
	StderrRedirect string
}
