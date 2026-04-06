package deeperr

import (
	"fmt"
)

type errorImpl struct {
	code    Code
	msg     string
	file    string
	line    int
	prevErr error
}

func (e errorImpl) Code() Code {
	return e.code
}

func (e errorImpl) Location() (string, int) {
	return e.file, e.line
}

func (e errorImpl) Error() string {
	if e.code != CodeUnset {
		return fmt.Sprintf("E%d %s", e.code, e.msg)
	}

	return e.msg
}

func (e errorImpl) Unwrap() error {
	return e.prevErr
}
