package main

import (
	"os"

	"github.com/sanity-io/litter"
	"github.com/thutasann/go-parser/src/lexer"
	"github.com/thutasann/go-parser/src/parser"
)

func main() {
	bytes, _ := os.ReadFile("./examples/04.lang")
	tokens := lexer.Tokenize(string(bytes))

	ast := parser.Parse(tokens)
	litter.Dump(ast)
}
