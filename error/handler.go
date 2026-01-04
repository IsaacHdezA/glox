package error

import (
	"fmt"
	"os"
)

type LoxError struct {
	line    int
	where   string
	message string
}

func NewLoxError(line int, where string, message string) *LoxError {
	loxError := new(LoxError)

	loxError.line = line
	loxError.where = where
	loxError.message = message

	return loxError
}

func (e *LoxError) Report() {
	fmt.Fprintf(os.Stderr, "[line %d] Error%s: %s\n", e.line, e.where, e.message)
}
