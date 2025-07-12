package lexer

import (
	"fmt"
	"regexp"
)

// Regex handler function type
type regexHandler func(lex *lexer, regex *regexp.Regexp)

// Regex pattern struct
type regexPattern struct {
	regex   *regexp.Regexp
	handler regexHandler
}

// Lexer struct
// Example:
//
// - createLexer("12 + 34")
// - It goes through each pattern and tries to match the start of the string.
//   - Matches 12 as a number → calls numberHandler → adds NUMBER token.
//   - Matches whitespace → calls skipHandler → skips.
//   - Matches + → calls defaultHandler(PLUS, "+") → adds PLUS token.
//   - Matches 34 → again NUMBER.
type lexer struct {
	patterns []regexPattern
	Tokens   []Token
	source   string
	pos      int
}

// Tokenize the source string
func Tokenize(source string) []Token {
	lex := createLexer(source)

	// 10 + [5]
	// Iterate while we sill have tokens
	for !lex.at_eof() {
		matched := false

		for _, pattern := range lex.patterns {
			loc := pattern.regex.FindStringIndex(lex.remainder())

			if loc != nil && loc[0] == 0 {
				pattern.handler(lex, pattern.regex)
				matched = true
				break
			}
		}

		// You could extend this to print the location and other things
		if !matched {
			panic(fmt.Sprintf("Lexer::Error -> unrecognized token near %s\n", lex.remainder()))
		}
	}

	lex.push(NewToken(EOF, "EOF"))
	return lex.Tokens
}

// Advance the position by n characters
// Finds a number match in the string, creates a token, and advances the position accordingly.
func (lex *lexer) advanceN(n int) {
	lex.pos += n
}

// Push a token to the tokens slice
func (lex *lexer) push(token Token) {
	lex.Tokens = append(lex.Tokens, token)
}

// Get the character at the current position
func (lex *lexer) at() byte {
	return lex.source[lex.pos]
}

// Check if the current position is at the end of the source string
func (lex *lexer) at_eof() bool {
	return lex.pos >= len(lex.source)
}

// Get the remainder of the source string
func (lex *lexer) remainder() string {
	return lex.source[lex.pos:]
}

// Default handler for regex patterns
func defaultHandler(kind TokenKind, value string) regexHandler {
	return func(lex *lexer, regex *regexp.Regexp) {
		lex.advanceN(len(value))
		lex.push(NewToken(kind, value))
	}
}

// Skip handler for regex patterns
// Used for whitespace: skips over the matched part without generating a token.
func skipHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	lex.advanceN(match[1])
}

// Number handler for regex patterns
func numberHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())
	lex.push(NewToken(NUMBER, match))
	lex.advanceN(len(match))
}

// Create a new lexer
func createLexer(source string) *lexer {
	return &lexer{
		pos:    0,
		source: source,
		Tokens: make([]Token, 0),
		patterns: []regexPattern{
			{regexp.MustCompile(`[0-9]+(\.[0-9]+)?`), numberHandler},
			{regexp.MustCompile(`\s+`), skipHandler},
			{regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
			{regexp.MustCompile(`\]`), defaultHandler(CLOSE_BRACKET, "]")},
			{regexp.MustCompile(`\{`), defaultHandler(OPEN_CURLY, "{")},
			{regexp.MustCompile(`\}`), defaultHandler(CLOSE_CURLY, "}")},
			{regexp.MustCompile(`\(`), defaultHandler(OPEN_PAREN, "(")},
			{regexp.MustCompile(`\)`), defaultHandler(CLOSE_PAREN, ")")},
			{regexp.MustCompile(`==`), defaultHandler(EQUALS, "==")},
			{regexp.MustCompile(`!=`), defaultHandler(NOT_EQUALS, "!=")},
			{regexp.MustCompile(`=`), defaultHandler(ASSIGNMENT, "=")},
			{regexp.MustCompile(`!`), defaultHandler(NOT, "!")},
			{regexp.MustCompile(`<=`), defaultHandler(LESS_EQUALS, "<=")},
			{regexp.MustCompile(`<`), defaultHandler(LESS, "<")},
			{regexp.MustCompile(`>=`), defaultHandler(GREATER_EQUALS, ">=")},
			{regexp.MustCompile(`>`), defaultHandler(GREATER, ">")},
			{regexp.MustCompile(`\|\|`), defaultHandler(OR, "||")},
			{regexp.MustCompile(`&&`), defaultHandler(AND, "&&")},
			{regexp.MustCompile(`\.\.`), defaultHandler(DOT_DOT, "..")},
			{regexp.MustCompile(`\.`), defaultHandler(DOT, ".")},
			{regexp.MustCompile(`;`), defaultHandler(SEMI_COLON, ";")},
			{regexp.MustCompile(`:`), defaultHandler(COLON, ":")},
			{regexp.MustCompile(`\?`), defaultHandler(QUESTION, "?")},
			{regexp.MustCompile(`,`), defaultHandler(COMMA, ",")},
			{regexp.MustCompile(`\+\+`), defaultHandler(PLUS_PLUS, "++")},
			{regexp.MustCompile(`--`), defaultHandler(MINUS_MINUS, "--")},
			{regexp.MustCompile(`\+=`), defaultHandler(PLUS_EQUALS, "+=")},
			{regexp.MustCompile(`-=`), defaultHandler(MINUS_EQUALS, "-=")},
			{regexp.MustCompile(`\+`), defaultHandler(PLUS, "+")},
			{regexp.MustCompile(`-`), defaultHandler(DASH, "-")},
			{regexp.MustCompile(`/`), defaultHandler(SLASH, "/")},
			{regexp.MustCompile(`\*`), defaultHandler(STAR, "*")},
			{regexp.MustCompile(`%`), defaultHandler(PERCENT, "%")},
		},
	}
}
