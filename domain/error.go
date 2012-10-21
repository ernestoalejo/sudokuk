package domain

import (
	"fmt"
	"runtime/debug"
)

type Error struct {
	Original error
	Stack    []byte
}

func NewError(err error) error {
	return &Error{
		Original: err,
		Stack:    debug.Stack(),
	}
}

func NewErrorf(format string, args ...interface{}) error {
	return NewError(fmt.Errorf(format, args...))
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s\n\n%s", e.Original, e.Stack)
}
