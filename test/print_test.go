package test

import (
	"testing"

	"github.com/IsaacHdezA/glox/ast"
	"github.com/IsaacHdezA/glox/token"
)

func TestPrint(t *testing.T) {
	expr := ast.BinaryExpr{
		Left: ast.UnaryExpr{
			Operator: token.Token{Type: token.MINUS, Lexeme: "-"},
			Right:    ast.LiteralExpr{Value: 123},
		},
		Operator: token.Token{Type: token.STAR, Lexeme: "*"},
		Right: ast.GroupingExpr{
			Expression: ast.LiteralExpr{Value: 45.67},
		},
	}

	printer := ast.AstPrinter{}
	output := printer.Print(expr)
	correct := "(* (- 123) (group 45.67))"

	if output != correct {
		t.Fatalf("Output %q failed. Expecting %q", output, correct)
	}
}

func TestPrintRPN(t *testing.T) {
	expr := ast.BinaryExpr{
		Left: ast.GroupingExpr{
			Expression: ast.BinaryExpr{
				Left:     ast.LiteralExpr{Value: 1},
				Operator: token.Token{Type: token.PLUS, Lexeme: "+"},
				Right:    ast.LiteralExpr{Value: 2},
			},
		},
		Operator: token.Token{Type: token.STAR, Lexeme: "*"},
		Right: ast.GroupingExpr{
			Expression: ast.BinaryExpr{
				Left:     ast.LiteralExpr{Value: 4},
				Operator: token.Token{Type: token.MINUS, Lexeme: "-"},
				Right:    ast.LiteralExpr{Value: 3},
			},
		},
	}

	printer := ast.AstPrinterRPN{}
	output := printer.Print(expr)
	correct := "1 2 + 4 3 - *"

	if output != correct {
		t.Fatalf("Output %q failed. Expecting %q", output, correct)
	}
}
