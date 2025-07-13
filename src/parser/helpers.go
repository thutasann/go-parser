package parser

import "github.com/thutasann/go-parser/src/lexer"

// Returns the token at the current position without advancing
func (p *parser) currentToken() lexer.Token {
	return p.tokens[p.pos]
}

// Returns the kind/type of the current token, e.g., lexer.LET, lexer.NUMBER, etc.
func (p *parser) currentTokenKind() lexer.TokenKind {
	return p.currentToken().Kind
}

// Moves to the next token and returns the previous/current one before advancing
func (p *parser) advance() lexer.Token {
	tk := p.currentToken()
	p.pos++
	return tk
}

// Returns true if:
//
// - Parser hasn’t reached the end of the token list
//
// - Current token is not the special EOF (End of File) token
func (p *parser) hasTokens() bool {
	return p.pos < len(p.tokens) && p.currentTokenKind() != lexer.EOF
}
