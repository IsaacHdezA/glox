package error

import (
	"bufio"
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
	writer := bufio.NewWriter(os.Stderr)

	out := fmt.Sprintf("[line %d] Error %s: %s", e.line, e.where, e.message)
	writer.WriteString(out)
}
