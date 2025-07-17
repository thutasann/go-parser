package ast

import "github.com/thutasann/go-parser/src/lexer"

// -------------------
// LITERAL EXPRESSIONS
// -------------------

type NumberExpr struct {
	Value float64
}

func (n NumberExpr) expr() {

}

type StringExpr struct {
	Value string
}

func (n StringExpr) expr() {

}

type SymbolExpr struct {
	Value string
}

func (n SymbolExpr) expr() {
}

// -------------------
// COMPLEX EXPRESSIONS
// -------------------

type BinaryExpr struct {
	Left     Expr
	Operator lexer.Token
	Right    Expr
}

func (n BinaryExpr) expr() {

}

type PrefixExpr struct {
	Operator  lexer.Token
	RightExpr Expr
}

func (n PrefixExpr) expr() {
}

type AssignmentExpr struct {
	Assigne  Expr
	Operator lexer.Token
	Value    Expr
}

func (n AssignmentExpr) expr() {
}
