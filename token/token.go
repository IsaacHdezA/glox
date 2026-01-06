package token

import (
	"fmt"
	"strings"

	"github.com/IsaacHdezA/glox/common"
)

type Token struct {
	Lexeme  string
	Type    TokenType
	Literal any

	location *common.Location
}

func NewToken(Type TokenType, Lexeme string, location *common.Location, Literal any) *Token {
	token := new(Token)

	token.Type = Type
	token.Lexeme = Lexeme
	token.location = location
	token.Literal = Literal

	return token
}

func (t *Token) String() string {
	output := fmt.Sprintf("|%s", t.Type)

	switch t.Type {
	case NUMBER:
		output += fmt.Sprintf(" %f", t.Literal)

	case COMMENT:
		commentText := strings.Trim(t.Lexeme[2:], " \n\r\t")
		output += fmt.Sprintf(" %q", commentText)

	case MULTI_COMMENT:
		n := len(t.Lexeme)

		commentText := strings.Trim(t.Lexeme[2:n-2], " \n\r\t")
		output += fmt.Sprintf(" %q", commentText)

	case STRING:
		output += fmt.Sprintf(" %q", t.Literal)

	default:
		if t.Lexeme == "" {
			output += "|"
			return output
		}

		output += fmt.Sprintf(" %s", t.Lexeme)
	}
	output += "|"

	return output
}
