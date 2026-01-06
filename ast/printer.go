package ast

import (
	"fmt"
)

// Pretty Printer
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

// Reverse Polish Notation Printer
type AstPrinterRPN struct{}

func (p AstPrinterRPN) VisitBinaryExpr(expr BinaryExpr) any {

	return fmt.Sprintf("%s %s %s", expr.Left.Accept(p), expr.Right.Accept(p), expr.Operator.Lexeme)
}

func (p AstPrinterRPN) VisitGroupingExpr(expr GroupingExpr) any {
	return fmt.Sprintf("%s", expr.Expression.Accept(p))
}

func (p AstPrinterRPN) VisitLiteralExpr(expr LiteralExpr) any {
	return fmt.Sprintf("%v", expr.Value)
}

func (p AstPrinterRPN) VisitUnaryExpr(expr UnaryExpr) any {
	return fmt.Sprintf("%v%s", expr.Right.Accept(p), expr.Operator.Lexeme)
}

func (p AstPrinterRPN) Print(expr Expr) string {
	return expr.Accept(p).(string)
}
