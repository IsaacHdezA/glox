package scanner

import (
	"fmt"
	"strconv"

	"github.com/IsaacHdezA/glox/common"
	"github.com/IsaacHdezA/glox/error"
	"github.com/IsaacHdezA/glox/token"
)

type Scanner struct {
	source string
	tokens []*token.Token

	line    int
	start   int
	current int
}

func NewScanner(source string) *Scanner {
	scanner := new(Scanner)
	scanner.source = source

	return scanner
}

func (s *Scanner) ScanTokens(source string) ([]*token.Token, *error.LoxError) {
	for !s.isAtEnd() {
		s.start = s.current

		err := s.scanToken()
		if err != nil {
			return nil, err
		}
	}

	return s.tokens, nil
}

func (s *Scanner) scanToken() *error.LoxError {
	c := s.advance()

	switch c {
	case '(':
		s.addToken(token.LEFT_PAREN)

	case ')':
		s.addToken(token.RIGHT_PAREN)

	case '{':
		s.addToken(token.LEFT_BRACE)

	case '}':
		s.addToken(token.RIGHT_BRACE)

	case ',':
		s.addToken(token.COMMA)

	case '.':
		s.addToken(token.DOT)

	case '-':
		s.addToken(token.MINUS)

	case '+':
		s.addToken(token.PLUS)

	case ';':
		s.addToken(token.SEMICOLON)

	case '*':
		s.addToken(token.STAR)

	case '!':
		if s.match('=') {
			s.addToken(token.BANG_EQUAL)
		} else {
			s.addToken(token.BANG)
		}

	case '=':
		if s.match('=') {
			s.addToken(token.EQUAL_EQUAL)
		} else {
			s.addToken(token.EQUAL)
		}

	case '<':
		if s.match('=') {
			s.addToken(token.LESS_EQUAL)
		} else {
			s.addToken(token.LESS)
		}

	case '>':
		if s.match('=') {
			s.addToken(token.GREATER_EQUAL)
		} else {
			s.addToken(token.GREATER)
		}

	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(token.SLASH)
		}

	case '"':
		s.string()

	case '\n':
		s.line++

	case ' ':
	case '\r':
	case '\t':

	default:
		if s.isDigit(c) {
			s.number()
		} else {
			error.NewLoxError(s.line, "", fmt.Sprintf("Unexpected character: %q.", c)).Report()
		}
	}

	return nil
}

func (s *Scanner) isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (s *Scanner) string() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}

		s.advance()
	}

	if s.isAtEnd() {
		str := s.source[s.start:s.current]
		error.NewLoxError(s.line, "", fmt.Sprintf("Unterminated string: %q", str)).Report()
	}

	// The closing quote (")
	s.advance()
	s.addToken(token.STRING)

}

func (s *Scanner) number() {
	for s.isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		s.advance()

		for s.isDigit(s.peek()) {
			s.advance()
		}
	}

	lexeme := s.source[s.start:s.current]
	number, _ := strconv.ParseFloat(lexeme, 64)

	s.addTokenLiteral(token.NUMBER, number)
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}

	if s.source[s.current] != expected {
		return false
	}

	s.current++

	return true
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	}

	return s.source[s.current]
}

func (s *Scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return 0
	}

	return s.source[s.current+1]
}

func (s *Scanner) advance() byte {
	c := s.source[s.current]
	s.current++

	return c
}

func (s *Scanner) addToken(_type token.TokenType) {
	lexeme := s.source[s.start:s.current]

	loc := common.NewLocation(s.line, 0, s.start, len(lexeme))
	token := token.NewToken(_type, lexeme, loc, 0)

	s.tokens = append(s.tokens, token)
}

func (s *Scanner) addTokenLiteral(_type token.TokenType, literal float64) {
	lexeme := s.source[s.start:s.current]

	loc := common.NewLocation(s.line, 0, s.start, len(lexeme))
	token := token.NewToken(_type, lexeme, loc, literal)

	s.tokens = append(s.tokens, token)
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}
