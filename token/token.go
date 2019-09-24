package token

import "fmt"

// Type is an interface that you custom token types must implement in order to
// be used with this lexer.
type Type interface {
	Name() string
	Value() uint64
}

// Token is a struct that holds a token type (defined by you), a value (the
// runes that were consumed by the lexer as string) and a position (where the
// token starts).
type Token struct {
	Type  Type
	Value string
	Pos   int
}

// New creates a new Token with the given Type, Value and Position.
func New(t Type, v string, pos int) Token {
	return Token{
		Type:  t,
		Value: v,
		Pos:   pos,
	}
}

// String returns a string representation of the form
//
//	Type.Name() + "(" + Value + "), " + Pos
//
// For example
//
//	Number(7.5), pos 9
func (t Token) String() string {
	return fmt.Sprintf("%v(%v), pos %d", t.Type.Name(), t.Value, t.Pos)
}
