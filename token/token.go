package token

import "fmt"

type Type interface {
	Name() string
	Value() uint64
}

type Token struct {
	Type  Type
	Value string
	Pos   int
}

func New(t Type, v string, pos int) Token {
	return Token{
		Type:  t,
		Value: v,
		Pos:   pos,
	}
}

func (t Token) String() string {
	return fmt.Sprintf("%v(%v), pos %d", t.Type.Name(), t.Value, t.Pos)
}
