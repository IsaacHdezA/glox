package ast

import (
	"fmt"
)

type AstPrinter struct{}

func (p AstPrinter) VisitBinaryExpr(expr BinaryExpr) any {
	return p.parentezize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (p AstPrinter) VisitGroupingExpr(expr GroupingExpr) any {
	return p.parentezize("group", expr.Expression)
}

func (p AstPrinter) VisitLiteralExpr(expr LiteralExpr) any {
	if expr.Value == nil {
		return "nil"
	}

	return fmt.Sprintf("%v", expr.Value)
}

func (p AstPrinter) VisitUnaryExpr(expr UnaryExpr) any {
	return p.parentezize(expr.Operator.Lexeme, expr.Right)
}

func (p AstPrinter) Print(expr Expr) string {
	return expr.Accept(p).(string)
}

func (p AstPrinter) parentezize(name string, exprs ...Expr) string {
	output := fmt.Sprintf("(%s", name)

	for _, expr := range exprs {
		output += " "
		output += expr.Accept(p).(string)
	}

	output += ")"

	return output
}
