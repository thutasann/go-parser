package ast

// Statement Interface
type Stmt interface {
	stmt()
}

// Expression Interface
type Expr interface {
	expr()
}

type Type interface {
	_type()
}
