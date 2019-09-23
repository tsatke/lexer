package lexer

import (
	"unicode/utf8"

	"github.com/TimSatke/parser/token"
)

var _ Lexer = (*baseLexer)(nil)

type baseLexer struct {
	input []byte
	start int
	pos   int
	width int

	current State
	tokens  token.Stream
}

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

func (l *baseLexer) TokenStream() token.Stream {
	return l.tokens
}

func (l *baseLexer) Emit(t token.Type) {
	l.tokens.Push(token.New(t, string(l.input[l.start:l.pos]), l.start))
	l.start = l.pos
}

func (l *baseLexer) EmitError(t token.Type, msg string) {
	l.tokens.Push(token.New(t, msg, l.pos))
	l.start = l.pos
}

func (l *baseLexer) Peek() rune {
	r := l.Next()
	l.Backup()
	return r
}

func (l *baseLexer) IsEOF() bool {
	return l.pos >= len(l.input)
}

func (l *baseLexer) Next() rune {
	r, width := utf8.DecodeRune(l.input[l.pos:])
	l.width = width
	l.pos += width

	return r
}

func (l *baseLexer) Ignore() {
	l.start = l.pos
}

func (l *baseLexer) Backup() {
	l.pos -= l.width
}

func (l *baseLexer) Accept(cc CharacterClass) bool {
	if cc.Matches(l.Next()) {
		return true
	}
	l.Backup()
	return false
}

func (l *baseLexer) AcceptMultiple(cc CharacterClass) bool {
	var matched bool
	for cc.Matches(l.Next()) {
		matched = true
	}
	l.Backup()
	return matched
}
