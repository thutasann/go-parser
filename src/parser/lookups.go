package parser

import (
	"github.com/thutasann/go-parser/src/ast"
	"github.com/thutasann/go-parser/src/lexer"
)

// Binding Power — defines operator precedence
type binding_power int

const (
	default_bp     binding_power = iota
	comma                        // ,
	assignment                   // =, +=, -=, etc.
	logical                      // &&, ||, ..
	relational                   // ==, !=, <, >, <=, >=
	additive                     // +, -
	multiplicative               // *, /, %
	unary                        // -, ! (prefix ops)
	call                         // ()
	member                       // .
	primary                      // literals, identifiers
)

// Statement Handler - parses a statement
type stmt_handler func(p *parser) ast.Stmt

// Nud handler — parses prefix expressions or literals
type nud_handler func(p *parser) ast.Expr

// Led handler — parses infix/postfix expressions
type led_handler func(p *parser, left ast.Expr, bp binding_power) ast.Expr

// Statement Lookup table for tokens
type stmt_lookup map[lexer.TokenKind]stmt_handler

// Null Denotation Lookup table for tokens
type nud_lookup map[lexer.TokenKind]nud_handler

// Left Denotation Lookup table for tokens
type led_lookup map[lexer.TokenKind]led_handler

// Binding Power Lookup table for tokens
type bp_lookup map[lexer.TokenKind]binding_power

var (
	bp_lu   = bp_lookup{}   // Token → binding power
	nud_lu  = nud_lookup{}  // Token → nud handler
	led_lu  = led_lookup{}  // Token → led handler
	stmt_lu = stmt_lookup{} // Token → statement handler
)

// Register a left denotation (infix/postfix) handler
func led(kind lexer.TokenKind, bp binding_power, led_fn led_handler) {
	bp_lu[kind] = bp
	led_lu[kind] = led_fn
}

// Register a null denotation (literal/prefix) handler
func nud(kind lexer.TokenKind, nud_fn nud_handler) {
	nud_lu[kind] = nud_fn
}

// Register a statement handler
func stmt(kind lexer.TokenKind, stmt_fn stmt_handler) {
	bp_lu[kind] = default_bp
	stmt_lu[kind] = stmt_fn
}

// Initializes all token lookups with appropriate handlers and precedence
func createTokenLookups() {
	led(lexer.ASSIGNMENT, assignment, parse_assignment_expr)
	led(lexer.PLUS_EQUALS, assignment, parse_assignment_expr)
	led(lexer.MINUS_EQUALS, assignment, parse_assignment_expr)

	// Logical
	led(lexer.AND, logical, parse_binary_expr)
	led(lexer.OR, logical, parse_binary_expr)
	led(lexer.DOT_DOT, logical, parse_binary_expr)

	// Relational
	led(lexer.LESS, relational, parse_binary_expr)
	led(lexer.LESS_EQUALS, relational, parse_binary_expr)
	led(lexer.GREATER, relational, parse_binary_expr)
	led(lexer.GREATER_EQUALS, relational, parse_binary_expr)
	led(lexer.EQUALS, relational, parse_binary_expr)
	led(lexer.NOT_EQUALS, relational, parse_binary_expr)

	// Additive & Multiplicative
	led(lexer.PLUS, additive, parse_binary_expr)
	led(lexer.DASH, additive, parse_binary_expr)
	led(lexer.STAR, multiplicative, parse_binary_expr)
	led(lexer.SLASH, multiplicative, parse_binary_expr)
	led(lexer.PERCENT, multiplicative, parse_binary_expr)

	// Literals & symbols
	nud(lexer.NUMBER, parse_primary_expr)
	nud(lexer.STRING, parse_primary_expr)
	nud(lexer.IDENTIFIER, parse_primary_expr)

	nud(lexer.DASH, parse_prefix_expr)

	// Statements
	stmt(lexer.CONST, parse_var_decl_stmt)
	stmt(lexer.LET, parse_var_decl_stmt)
}
