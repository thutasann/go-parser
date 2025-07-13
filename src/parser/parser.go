package parser

import (
	"github.com/thutasann/go-parser/src/ast"
	"github.com/thutasann/go-parser/src/lexer"
)

// Holds all the tokens from the lexer
// pos: current position/index in the token list
type parser struct {
	tokens []lexer.Token
	pos    int
}

// Creates a parser instance.
//
// - Calls createTokenLookups() → this registers all the nud/led/stmt handlers and operator precedence.
//
// - Initializes pos to 0.
func createParser(tokens []lexer.Token) *parser {
	createTokenLookups()
	return &parser{
		tokens: tokens,
		pos:    0,
	}
}

// - Function to parse tokens into an `ast.BlockStmt`.
//
// - Creates the parser object
//
// - Parses statements in a loop until all tokens are consumed
//
// - Adds each statement into Body
//
// - Returns a block statement, which wraps all the parsed statements
func Parse(tokens []lexer.Token) ast.BlockStmt {
	Body := make([]ast.Stmt, 0)
	p := createParser(tokens)

	for p.hasTokens() {
		Body = append(Body, parse_stmt(p))
	}

	return ast.BlockStmt{
		Body: Body,
	}
}
