package ast

// Block Statement { ... []Stmt }
type BlockStmt struct {
	Body []Stmt
}

func (n BlockStmt) stmt() {}

// Expression Statement
type ExpressionStmt struct {
	Expression Expr
}

func (n ExpressionStmt) stmt() {}

// Variable Declaration Statement
type VarDeclStmt struct {
	VariableName  string
	IsConstant    bool
	AssignedValue Expr
	ExplicitType  Type
}

func (n VarDeclStmt) stmt() {}
