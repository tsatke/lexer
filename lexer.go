package lexer

import (
	"github.com/TimSatke/lexer/token"
)

// Lexer is an interface providing all nececssary methods for lexing text. There
// is a default implementation for UTF-8 input, which can be used as follows.
//
//	func main() {
//		l := lexer.New(input, lexRoot)
//		_ = l
//	}
//
// See the examples in the godoc for more information.
type Lexer interface {
	// StartLexing will cause the lexer to start pushing tokens onto the token
	// stream. See the documentation of the implementing struct for information
	// on how to use this.
	StartLexing()

	// TokenStream returns the token stream that the lexer will push tokens
	// onto.
	TokenStream() token.Stream
	// Emit pushes a token of the given type with its position and all consumes
	// runes onto the token stream.
	Emit(token.Type)
	// EmitError emits an error token with the given error token type (that was
	// defined by you) and a given error message.
	EmitError(token.Type, string)
	// IsEOF determines whether the lexer has already reached the end of the
	// input.
	IsEOF() bool
	// Peek reads the next rune, but does not consume it. Peek does not advance
	// the lexer position in the input.
	Peek() rune
	// Next reads the next rune and consumes it. Next advances the lexer
	// position in the input by the byte-width of the read rune.
	Next() rune
	// Ignore discards all consumes runes. This behaves like Emit(...), except
	// it doesn't create/push a token onto the token stream.
	Ignore()
	// Backup unreads the last consumed rune.
	Backup()

	// Accept consumes the next rune, if and only if it is matched by the given
	// character class. Accept returns true if the next rune was matched and
	// consumed.
	Accept(CharacterClass) bool
	// AcceptMultiple consumes the next N runes that are matched by the given
	// character class. AcceptMultiple returns the amount of runes that were
	// matched.
	AcceptMultiple(CharacterClass) uint
}
