package deeperr

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// GetStacktrace traverses the error chain and constructs a formatted string, representing the full execution path.
// for each error in the chain, it prints the error message and if the error is a deeperr.Error it will
// also print the location where the error was created.
func GetStacktrace(err error) string {
	var sb strings.Builder

	const indent = 4

	curr := err
	for curr != nil {
		fmt.Fprintf(&sb, "%s\n", curr.Error())

		if e, ok := curr.(errorImpl); ok && e.file != "" { // nolint:errorlint
			fmt.Fprintf(&sb, "%s%s:%d\n", strings.Repeat(" ", indent), e.file, e.line)
		}

		curr = errors.Unwrap(curr)
	}

	return sb.String()
}

// PrintStacktrace is a helper function that retrieves the full stack trace of an error and writes it directly to
// standard error (os.Stderr).
func PrintStacktrace(err error) {
	fmt.Fprintln(os.Stderr, GetStacktrace(err))
}
