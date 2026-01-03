package scanner

import (
	"github.com/IsaacHdezA/glox/error"
	"github.com/IsaacHdezA/glox/token"
)

type Scanner struct {
	source   string
	hasError bool
}

func NewScanner(source string) *Scanner {
	scanner := new(Scanner)
	scanner.source = source

	return scanner
}

func (*Scanner) ScanTokens(source string) ([]*token.Token, *error.LoxError) {
	tokens := []*token.Token{}

	return tokens, nil
}

func (s *Scanner) HasError() bool {
	return s.hasError
}

func (s *Scanner) SetError(hasError bool) {
	s.hasError = hasError
}
