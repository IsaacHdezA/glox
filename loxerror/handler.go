package loxerror

import (
	"fmt"
	"github.com/IsaacHdezA/glox/token"
)

type ScannerError struct {
	Line    int
	Where   string
	Message string
}

func NewScannerError(Line int, Where string, Message string) *ScannerError {
	loxError := new(ScannerError)

	loxError.Line = Line
	loxError.Where = Where
	loxError.Message = Message

	return loxError
}

func (e *ScannerError) Error() string {
	return fmt.Sprintf("[line %d] Error%s: %s", e.Line, e.Where, e.Message)
}

type ParserError struct {
	Token   token.Token
	Message string
}

func NewParserError(
	Token token.Token,
	Message string,
) *ParserError {
	loxError := new(ParserError)

	loxError.Token = Token
	loxError.Message = Message

	return loxError
}

func (e *ParserError) Error() string {
	if e.Token.Type == token.EOF {
		return fmt.Sprintf("%d at end %s", e.Token.Location.Line, e.Message)
	}

	return fmt.Sprintf("%d at %q %s", e.Token.Location.Line, e.Token.Lexeme, e.Message)
}
