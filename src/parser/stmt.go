package parser

import (
	"github.com/thutasann/go-parser/src/ast"
	"github.com/thutasann/go-parser/src/lexer"
)

// Parse Statement
func parse_stmt(p *parser) ast.Stmt {
	smt_fn, exists := stmt_lu[p.currentTokenKind()]

	if exists {
		return smt_fn(p)
	}

	expression := parse_expr(p, default_bp)
	p.expect(lexer.SEMI_COLON)

	return ast.ExpressionStmt{
		Expression: expression,
	}
}
