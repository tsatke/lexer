package main

import (
	"fmt"

	"github.com/TimSatke/parser/lexer"
)

const data = `
{
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

	for t := range l.TokenStream().Tokens() {
		fmt.Printf("Token: %v\n", t.String())
	}
}
