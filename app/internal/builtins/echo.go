package builtins

import (
	"fmt"
	"io"
	"strings"
)

func Echo(args []string, out io.Writer) {
	fmt.Fprintln(out, strings.Join(args, " "))
}
