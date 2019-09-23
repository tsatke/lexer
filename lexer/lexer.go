package lexer

import (
	"github.com/TimSatke/parser/token"
)

type State func(Lexer) State

type Lexer interface {
	StartLexing()

	TokenStream() token.Stream
	Emit(token.Type)
	EmitError(token.Type, string)
	IsEOF() bool
	Peek() rune
	Next() rune
	Ignore()
	Backup()

	Accept(CharacterClass) bool
	AcceptMultiple(CharacterClass) bool
}
