// Package deeperr is a lightweight Go error-wrapping library designed to provide stack traces and error codes
package deeperr

import (
	"errors"
	"runtime"
)

// Code represents an error code
type Code int

// CodeUnset means this error doesn't use a code
const CodeUnset Code = -1

// Error is an error type with stack traces and error code tagging capabilities
type Error interface {
	Code() Code
	Error() string
	Unwrap() error
	Location() (string, int)
}

// New returns a new Error
func New(message string, prev error) Error {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = ""
		line = -1
	}

	return errorImpl{
		code:    CodeUnset,
		msg:     message,
		file:    file,
		line:    line,
		prevErr: prev,
	}
}

// NewWithCode returns an Error with the associated code
func NewWithCode(code Code, message string, prev error) Error {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = ""
		line = -1
	}

	return errorImpl{
		code:    code,
		msg:     message,
		file:    file,
		line:    line,
		prevErr: prev,
	}
}

// IsCode checks if the first deeperr.Error in the chain has the error code
func IsCode(err error, code Code) bool {
	if e, ok := errors.AsType[Error](err); ok {
		return e.Code() == code
	}

	return false
}

// Contains checks the entire chain and returns true if ANY deeperr.Error in the hierarchy matches the error code
func Contains(err error, code Code) bool {
	for err != nil {
		if e, ok := err.(Error); ok { // nolint:errorlint
			if e.Code() == code {
				return true
			}
		}

		err = errors.Unwrap(err)
	}

	return false
}
