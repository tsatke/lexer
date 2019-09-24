package lexer

import (
	"unicode/utf8"

	"github.com/TimSatke/lexer/token"
)

var _ Lexer = (*baseLexer)(nil) // ensure that baseLexer implements Lexer interface

// baseLexer is a default implementation of a lexer. It is fully usable and if
// you're using it, you only have to define your lexer states.
type baseLexer struct {
	input []byte
	start int
	pos   int
	width int

	current State
	tokens  token.Stream
}

// New returns a usable lexer implementation, that takes an input which will be
// lexed and a start state, which has to be provided. This state will be
// executed until nil is returned as a next state.
func New(input []byte, start State) *baseLexer {
	return &baseLexer{
		input: input,
		start: 0,
		pos:   0,
		width: 0,

		current: start,
		tokens:  token.NewStream(),
	}
}

// StartLexing starts the lexing of the given input. This method is blocking, so
// it is recommended to execute it in a separate goroutine. It will stop once
// all bytes from the input data are read. When it's done lexing, it will close
// the token stream.
func (l *baseLexer) StartLexing() {
	defer l.tokens.Close()

	for !l.IsEOF() {
		l.current = l.current(l)
		if l.current == nil {
			// last state was end state
			break
		}
	}
}

// TokenStream returns the token stream that all tokens will be pushed onto
// while lexing.
//
//	l := lexer.New(...)
//	for token := range l.TokenStream().Tokens() {
//		fmt.Println(token.String)
//	}
func (l *baseLexer) TokenStream() token.Stream {
	return l.tokens
}

// Emit emits all accepted runes since the last emit as a token. This means,
// that a new token with the accepted runes and the given type is created and
// pushed onto the token stream.
func (l *baseLexer) Emit(t token.Type) {
	l.tokens.Push(token.New(t, string(l.input[l.start:l.pos]), l.start))
	l.start = l.pos
}

// EmitError emits a token with the given message and token type and pushes it
// onto the token stream. The token stream is required, since you have to define
// all token by yourself, and this lexer can neither know the error token
// type(s) you use, nor which ones your parser will handle how.
//
//	const (
//		MyErrorTokenType token.Type = iota
//		MyWhitespaceTokenType
//	)
//	func lexWhitespace(l lexer.Lexer) lexer.State {
//		if !l.Accept(CCWhitespace) {
//			return errorf("Unexpected token %v, expected one of %v", l.Peek(), CCWhitespace.String())
//		}
//		l.Emit(MyWhitespaceTokenType)
//		return nil
//	}
//	func errorf(format string, args ...interface{}) lexer.State {
//		return func(l lexer.Lexer) lexer.State {
//			l.EmitError(MyErrorTokenType, fmt.Sprintf(format, args...))
//			return nil // no next state; lexer will stop
//		}
//	}
func (l *baseLexer) EmitError(t token.Type, msg string) {
	l.tokens.Push(token.New(t, msg, l.pos))
	l.start = l.pos
}

// Peek returns the next rune that the lexer will read upon the next call of
// Next(). This does not consume the next rune.
func (l *baseLexer) Peek() rune {
	r := l.Next()
	l.Backup()
	return r
}

// IsEOF determines whether there are more bytes left to read.
func (l *baseLexer) IsEOF() bool {
	return l.pos >= len(l.input)
}

// Next consumes the next rune, advancing the lexers position by (rune-width).
// The rune read will be utf8 decoded from the input.
func (l *baseLexer) Next() rune {
	r, width := utf8.DecodeRune(l.input[l.pos:])
	l.width = width
	l.pos += width

	return r
}

// Ignore discards all accepted runes and updates the token start marker to the
// current position.
func (l *baseLexer) Ignore() {
	l.start = l.pos
}

// Backup unreads the last read rune.
func (l *baseLexer) Backup() {
	l.pos -= l.width
}

// Accept accepts one rune from the given character class. If the next rune is
// matched by the character class, it is consumed. Otherwise, it is not
// consumed. This method returns whether the rune was consumed or not, or, in
// other words, determines, whether the next rune is matched by the character
// class.
func (l *baseLexer) Accept(cc CharacterClass) bool {
	if cc.Matches(l.Next()) {
		return true
	}
	l.Backup()
	return false
}

// AcceptMultiple consumes the next N consecutive runes that are matched by the
// given character class. It returns the amount of runes that were consumed, or
// 0 if none were matched.
func (l *baseLexer) AcceptMultiple(cc CharacterClass) uint {
	var matched uint
	for cc.Matches(l.Next()) {
		matched++
	}
	l.Backup()
	return matched
}
