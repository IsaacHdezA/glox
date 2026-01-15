package parser

import (
	"fmt"
	"os"
	"slices"

	"github.com/IsaacHdezA/glox/ast"
	"github.com/IsaacHdezA/glox/loxerror"
	"github.com/IsaacHdezA/glox/token"
)

type Parser struct {
	Tokens []token.Token

	Current int
}

func NewParser(
	Tokens []token.Token,
) *Parser {
	parser := new(Parser)

	parser.Tokens = Tokens

	return parser
}

func (p *Parser) Parse() (ast.Expr, error) {
	expr, err := p.expression()

	if err != nil {
		return nil, err
	}

	return expr, nil
}

func (p *Parser) expression() (ast.Expr, error) {
	return p.equality()
}

func (p *Parser) equality() (ast.Expr, error) {
	expr, err := p.comparison()

	if err != nil {
		return nil, err
	}

	for p.match(token.BANG_EQUAL, token.EQUAL_EQUAL) {
		operator := p.previous()
		right, err := p.comparison()

		if err != nil {
			return nil, err
		}

		expr = ast.BinaryExpr{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *Parser) match(tokens ...token.TokenType) bool {
	return slices.ContainsFunc(tokens, func(token token.TokenType) bool {
		checks := p.check(token)

		if checks == true {
			p.advance()

			return checks
		}

		return checks
	})
}

func (p *Parser) check(token token.TokenType) bool {
	if p.isAtEnd() {
		return false
	}

	return p.peek().Type == token
}

func (p *Parser) advance() token.Token {
	if !p.isAtEnd() {
		p.Current++
	}

	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == token.EOF
}

func (p *Parser) peek() token.Token {
	return p.Tokens[p.Current]
}

func (p *Parser) previous() token.Token {
	return p.Tokens[p.Current-1]
}

func (p *Parser) error(token token.Token, msg string) *loxerror.ParserError {
	err := loxerror.NewParserError(token, msg)
	fmt.Fprintln(os.Stderr, err)

	return err
}

func (p *Parser) sinchronize() {
	p.advance()

	for !p.isAtEnd() {
		if p.previous().Type == token.SEMICOLON {
			return
		}

		switch p.peek().Type {
		case token.CLASS:
			return
		case token.FUN:
			return
		case token.VAR:
			return
		case token.FOR:
			return
		case token.IF:
			return
		case token.WHILE:
			return
		case token.PRINT:
			return
		case token.RETURN:
			return
		}

		p.advance()
	}
}

func (p *Parser) comparison() (ast.Expr, error) {
	expr, err := p.term()

	if err != nil {
		return nil, err
	}

	for p.match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		operator := p.previous()
		right, err := p.term()

		if err != nil {
			return nil, err
		}

		expr = ast.BinaryExpr{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *Parser) term() (ast.Expr, error) {
	expr, err := p.factor()

	if err != nil {
		return nil, err
	}

	for p.match(token.MINUS, token.PLUS) {
		operator := p.previous()
		right, err := p.factor()

		if err != nil {
			return nil, err
		}

		expr = ast.BinaryExpr{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *Parser) factor() (ast.Expr, error) {
	expr, err := p.unary()

	if err != nil {
		return nil, err
	}

	for p.match(token.SLASH, token.STAR) {
		operator := p.previous()
		right, err := p.unary()

		if err != nil {
			return nil, err
		}

		expr = ast.BinaryExpr{Left: expr, Operator: operator, Right: right}
	}

	return expr, nil
}

func (p *Parser) unary() (ast.Expr, error) {
	if p.match(token.BANG, token.MINUS) {
		operator := p.previous()
		right, err := p.unary()

		if err != nil {
			return nil, err
		}

		return ast.UnaryExpr{Operator: operator, Right: right}, nil
	}

	return p.primary()
}

func (p *Parser) primary() (ast.Expr, error) {
	if p.match(token.FALSE) {
		return ast.LiteralExpr{Value: false}, nil
	}
	if p.match(token.TRUE) {
		return ast.LiteralExpr{Value: true}, nil
	}
	if p.match(token.NIL) {
		return ast.LiteralExpr{Value: nil}, nil
	}

	if p.match(token.STRING, token.NUMBER) {
		return ast.LiteralExpr{Value: p.previous().Literal}, nil
	}

	if p.match(token.LEFT_PAREN) {
		expr, err := p.expression()

		if err != nil {
			return nil, err
		}

		p.consume(token.RIGHT_PAREN, "Expect ')' after expression")

		return ast.GroupingExpr{Expression: expr}, nil
	}

	return nil, p.error(p.peek(), "Expected expression.")
}

func (p *Parser) consume(token token.TokenType, errorMsg string) (token.Token, error) {
	if p.check(token) {
		return p.advance(), nil
	}

	return p.peek(), p.error(p.peek(), errorMsg)
}
