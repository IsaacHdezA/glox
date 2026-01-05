package ast

import "github.com/IsaacHdezA/glox/token"

type Expr interface {
	Accept(visitor *ExprVisitor) any
}

type ExprVisitor interface {
	VisitBinaryExpr(expr *BinaryExpr) any
	VisitGroupingExpr(expr *GroupingExpr) any
	VisitLiteralExpr(expr *LiteralExpr) any
	VisitUnaryExpr(expr *UnaryExpr) any
}

type BinaryExpr struct {
	left     Expr
	operator token.Token
	right    Expr
}

func NewBinaryExpr(
	left Expr,
	operator token.Token,
	right Expr,
) *BinaryExpr {
	result := new(BinaryExpr)

	result.left = left
	result.operator = operator
	result.right = right

	return result
}

func (e *BinaryExpr) Accept(visitor ExprVisitor) any {
	return visitor.VisitBinaryExpr(e)
}

type GroupingExpr struct {
	expression Expr
}

func NewGroupingExpr(
	expression Expr,
) *GroupingExpr {
	result := new(GroupingExpr)

	result.expression = expression

	return result
}

func (e *GroupingExpr) Accept(visitor ExprVisitor) any {
	return visitor.VisitGroupingExpr(e)
}

type LiteralExpr struct {
	value any
}

func NewLiteralExpr(
	value any,
) *LiteralExpr {
	result := new(LiteralExpr)

	result.value = value

	return result
}

func (e *LiteralExpr) Accept(visitor ExprVisitor) any {
	return visitor.VisitLiteralExpr(e)
}

type UnaryExpr struct {
	operator token.Token
	right    Expr
}

func NewUnaryExpr(
	operator token.Token,
	right Expr,
) *UnaryExpr {
	result := new(UnaryExpr)

	result.operator = operator
	result.right = right

	return result
}

func (e *UnaryExpr) Accept(visitor ExprVisitor) any {
	return visitor.VisitUnaryExpr(e)
}
