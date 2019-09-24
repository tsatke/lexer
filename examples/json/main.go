package main

import (
	"log"

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
	for token := range l.TokenStream().Tokens() {
		log.Printf("token: %v\n", token.String())
	}
}
