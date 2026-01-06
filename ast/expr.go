package ast

import "github.com/IsaacHdezA/glox/token"

type Expr interface {
	Accept(visitor ExprVisitor) any
}

type ExprVisitor interface {
	VisitBinaryExpr(expr BinaryExpr) any
	VisitGroupingExpr(expr GroupingExpr) any
	VisitLiteralExpr(expr LiteralExpr) any
	VisitUnaryExpr(expr UnaryExpr) any
}

type BinaryExpr struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

func NewBinaryExpr(
	Left Expr,
	Operator token.Token,
	Right Expr,
) BinaryExpr {
	result := BinaryExpr{}

	result.Left = Left
	result.Operator = Operator
	result.Right = Right

	return result
}

func (e BinaryExpr) Accept(visitor ExprVisitor) any {
	return visitor.VisitBinaryExpr(e)
}

type GroupingExpr struct {
	Expression Expr
}

func NewGroupingExpr(
	Expression Expr,
) GroupingExpr {
	result := GroupingExpr{}

	result.Expression = Expression

	return result
}

func (e GroupingExpr) Accept(visitor ExprVisitor) any {
	return visitor.VisitGroupingExpr(e)
}

type LiteralExpr struct {
	Value any
}

func NewLiteralExpr(
	Value any,
) LiteralExpr {
	result := LiteralExpr{}

	result.Value = Value

	return result
}

func (e LiteralExpr) Accept(visitor ExprVisitor) any {
	return visitor.VisitLiteralExpr(e)
}

type UnaryExpr struct {
	Operator token.Token
	Right    Expr
}

func NewUnaryExpr(
	Operator token.Token,
	Right Expr,
) UnaryExpr {
	result := UnaryExpr{}

	result.Operator = Operator
	result.Right = Right

	return result
}

func (e UnaryExpr) Accept(visitor ExprVisitor) any {
	return visitor.VisitUnaryExpr(e)
}
