package scanner

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/IsaacHdezA/glox/common"
	"github.com/IsaacHdezA/glox/loxerror"
	"github.com/IsaacHdezA/glox/token"
)

var _keywords = map[string]token.TokenType{
	"and":    token.AND,
	"class":  token.CLASS,
	"else":   token.ELSE,
	"false":  token.FALSE,
	"for":    token.FOR,
	"fun":    token.FUN,
	"if":     token.IF,
	"nil":    token.NIL,
	"or":     token.OR,
	"print":  token.PRINT,
	"return": token.RETURN,
	"super":  token.SUPER,
	"this":   token.THIS,
	"true":   token.TRUE,
	"var":    token.VAR,
	"while":  token.WHILE,
}

type Scanner struct {
	source string
	tokens []token.Token

	line    int
	start   int
	current int
}

func NewScanner(source string) *Scanner {
	scanner := new(Scanner)
	scanner.source = source

	return scanner
}

func (s *Scanner) ScanTokens(source string) ([]token.Token, *loxerror.ScannerError) {
	for !s.isAtEnd() {
		s.start = s.current

		err := s.scanToken()
		if err != nil {
			return nil, err
		}
	}

	s.start = s.current
	s.addToken(token.EOF)

	return s.tokens, nil
}

func (s *Scanner) scanToken() *loxerror.ScannerError {
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
		nextC := s.peek()
		if nextC == '/' || nextC == '*' {
			s.comment()
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
		} else if s.isAlpha(c) {
			s.identifier()
		} else {
			fmt.Fprintln(os.Stderr, loxerror.NewScannerError(s.line, "", fmt.Sprintf("Unexpected character: %q.", c)).Error())
		}
	}

	return nil
}

func (s *Scanner) identifier() {
	for s.isAlphanumeric(s.peek()) {
		s.advance()
	}

	lexeme := s.source[s.start:s.current]
	var tType token.TokenType
	tType, ok := _keywords[lexeme]

	if ok {
		s.addToken(tType)
	} else {
		s.addToken(token.IDENTIFIER)
	}
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
		fmt.Fprintln(os.Stderr, loxerror.NewScannerError(s.line, "", fmt.Sprintf("Unterminated string: %q", str)).Error())
	}

	// The closing quote (")
	s.advance()

	literal := s.source[s.start+1 : s.current-1]
	s.addTokenLiteral(token.STRING, literal)
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

func (s *Scanner) comment() {
	var text string

	if s.match('/') {
		s.singleComment()

		text = strings.Trim(s.source[s.start+2:s.current], " \n\t\r")
		if text != "" {
			s.addToken(token.COMMENT)
		}
	} else if s.match('*') {
		s.multiComment()

		n := s.current - s.start
		text = strings.Trim(s.source[s.start+2:s.start+(n-2)], " \n\t\r")
		if text != "" {
			s.addToken(token.MULTI_COMMENT)
		}
	}
}

func (s *Scanner) singleComment() {
	for s.peek() != '\n' && !s.isAtEnd() {
		s.advance()
	}
}

func (s *Scanner) multiComment() {
	if s.isAtEnd() {
		return
	}

	c1, c2 := s.peek(), s.peekNext()

	if c1 == '\n' {
		s.line++
	}

	if c1 == '*' && c2 == '/' {
		s.advance()
		s.advance()

		return
	}

	for !s.isAtEnd() {
		if c1 == '\n' {
			s.line++
		}

		c1, c2 = s.peek(), s.peekNext()

		if c1 == '/' && c2 == '*' {
			s.advance()
			s.advance()

			s.multiComment()
		} else if c1 == '*' && c2 == '/' {
			s.advance()
			s.advance()

			return
		} else if s.isAtEnd() {

			fmt.Fprintln(os.Stderr, loxerror.NewScannerError(s.line, "", "Unterminated multi-line comment").Error())
			return
		}

		s.advance()
	}
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

func (s *Scanner) isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func (s *Scanner) isAlphanumeric(c byte) bool {
	return s.isAlpha(c) || s.isDigit(c)
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

	s.tokens = append(s.tokens, *token)
}

func (s *Scanner) addTokenLiteral(_type token.TokenType, literal any) {
	lexeme := s.source[s.start:s.current]

	loc := common.NewLocation(s.line, 0, s.start, len(lexeme))
	token := token.NewToken(_type, lexeme, loc, literal)

	s.tokens = append(s.tokens, *token)
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}
