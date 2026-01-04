package token

import (
	"fmt"

	"github.com/IsaacHdezA/glox/common"
)

type Token struct {
	lexeme  string
	_type   TokenType
	literal float64

	location *common.Location
}

func NewToken(_type TokenType, lexeme string, location *common.Location, literal float64) *Token {
	token := new(Token)

	token._type = _type
	token.lexeme = lexeme
	token.location = location
	token.literal = literal

	return token
}

func (t *Token) String() string {
	output := fmt.Sprintf("%s %s", t._type, t.lexeme)

	switch t._type {
	case NUMBER:
		output += fmt.Sprintf(" %f", t.literal)
	case STRING:
		output += fmt.Sprintf(" %q", t.lexeme)
	}
	return output
}
