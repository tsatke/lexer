package main

type Token uint64

func (t Token) Name() string  { return tokenNames[t] }
func (t Token) Value() uint64 { return uint64(t) }

var (
	tokenNames = map[Token]string{
		TokenBraceOpen:    "BraceOpen",
		TokenBraceClose:   "BraceClose",
		TokenBracketOpen:  "BracketOpen",
		TokenBracketClose: "BracketClose",
		TokenColon:        "Colon",
		TokenComma:        "Comma",
		TokenError:        "Error",
		TokenString:       "String",
		TokenNumber:       "Number",
		TokenUnknown:      "Unknown",
		TokenWhitespace:   "Whitespace",

		TokenTrue:  "True",
		TokenFalse: "False",
		TokenNull:  "Null",
	}
)

const (
	TokenUnknown Token = iota
	TokenError

	TokenBraceOpen
	TokenBraceClose
	TokenBracketOpen
	TokenBracketClose
	TokenColon
	TokenComma
	TokenString
	TokenNumber
	TokenWhitespace

	TokenTrue
	TokenFalse
	TokenNull
)
