package main

import (
	"log"

	"github.com/TimSatke/parser/lexer"
	"github.com/TimSatke/parser/parser"
)

const data = `
{a
	"foo": "bar",
	"sna": true,
	"abc": [
		"d","e",false,
		"f",null,"g"
	],
	"random": [
		"\u00fd",
		"\\u00fd",
		"\u0000"
	],
	"numbers": [
		0e-5
	]
}
`

func main() {
	l := lexer.New([]byte(data), lexJSON)
	go l.StartLexing()

	p := parser.New(l, parseJSON, TokenUnknown, TokenError)

	err := p.Parse()
	if err != nil {
		log.Fatal(err)
	}
}

func parseJSON(p parser.Parser) parser.Rule {
	p.Accept(TokenWhitespace, TokenBracketOpen, TokenString)
	return parseJSON
}
