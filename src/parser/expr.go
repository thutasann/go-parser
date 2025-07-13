package parser

import (
	"fmt"
	"strconv"

	"github.com/thutasann/go-parser/src/ast"
	"github.com/thutasann/go-parser/src/lexer"
)

// Parse Primary Expression
func parse_primary_expr(p *parser) ast.Expr {
	switch p.currentTokenKind() {
	case lexer.NUMBER:
		number, _ := strconv.ParseFloat(p.advance().Value, 64)
		return ast.NumberExpr{
			Value: number,
		}
	case lexer.STRING:
		return ast.StringExpr{
			Value: p.advance().Value,
		}
	case lexer.IDENTIFIER:
		return ast.SymbolExpr{
			Value: p.advance().Value,
		}
	default:
		panic(fmt.Sprintf("cannot create primary_expression from %s\n", lexer.TokenKindString(p.currentTokenKind())))
	}
}

// Parse Binary Expression
func parse_binary_expr(p *parser, left ast.Expr, bp binding_power) ast.Expr {}
