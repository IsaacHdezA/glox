package token

import "fmt"

type Location struct {
	_sourceOffset int
	_tokenLength  int

	line   int
	column int
}

type Token struct {
	lexeme  string
	_type   TokenType
	literal float32

	location *Location
}

func NewToken(_type TokenType, lexeme string, location *Location, literal float32) *Token {
	token := new(Token)

	token._type = _type
	token.lexeme = lexeme
	token.location = location
	token.literal = literal

	return token
}

func (t *Token) String() string {
	output := fmt.Sprintf("[%s %s", t._type, t.lexeme)

	if t._type == NUMBER {
		output += fmt.Sprintf(" %f", t.literal)
	}

	output += "]"

	return output
}
