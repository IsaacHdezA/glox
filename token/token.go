package token

import (
	"fmt"
	"strings"

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
	output := fmt.Sprintf("%s", t._type)

	switch t._type {
	case NUMBER:
		output += fmt.Sprintf(" %f", t.literal)

	case COMMENT:
		commentText := strings.Trim(t.lexeme[2:], "\n \r\t")
		output += fmt.Sprintf(" %q", commentText)

	case MULTI_COMMENT:
		n := len(t.lexeme)
		commentText := strings.Trim(t.lexeme[2:n-2], "\n \r\t")

		output += fmt.Sprintf(" %q", commentText)

	case STRING:
		output += fmt.Sprintf(" %s", t.lexeme)

	default:
		output += fmt.Sprintf(" %s", t.lexeme)
	}
	return output
}
