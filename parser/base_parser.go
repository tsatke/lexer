package parser

import (
	"errors"
	"fmt"

	"github.com/TimSatke/parser/lexer"
	"github.com/TimSatke/parser/parser/ast"
	"github.com/TimSatke/parser/token"
)

type baseParser struct {
	stream      token.Stream
	errorTokens []token.Type
	tokens      []token.Token
	pos         int

	aborted bool
	err     chan error

	current Rule
}

func New(l lexer.Lexer, start Rule, errorTokens ...token.Type) *baseParser {
	return &baseParser{
		stream:      l.TokenStream(),
		pos:         0,
		errorTokens: errorTokens,

		aborted: false,
		err:     make(chan error),

		current: start,
	}
}

func (p *baseParser) Parse() error {
	errValues := make([]uint64, len(p.errorTokens))
	for i, et := range p.errorTokens {
		errValues[i] = et.Value()
	}

	for token := range p.stream.Tokens() {
		for _, v := range errValues {
			if v == token.Type.Value() {
				return fmt.Errorf("%v at pos %v", token.Value, token.Pos)
			}
		}
		p.tokens = append(p.tokens, token)
	}

	parsing := make(chan struct{})

	go func() {
		defer close(parsing)

		for p.HasMore() && !p.aborted {
			p.current = p.current(p)
			if p.current == nil {
				break
			}
		}
	}()

	select {
	case err := <-p.err:
		p.aborted = true
		return err
	case <-parsing:
		return nil
	}
}

func (p *baseParser) Tree() ast.Tree {
	return nil
}

func (p *baseParser) Error(msg string) {
	p.err <- errors.New(msg)
}

func (p *baseParser) HasMore() bool {
	return p.pos < len(p.tokens)
}

func (p *baseParser) Peek() (token.Token, bool) {
	if p.pos >= len(p.tokens)-1 {
		return token.Token{}, false
	}

	return p.tokens[p.pos+1], true
}

func (p *baseParser) Next() (token.Token, bool) {
	if p.pos >= len(p.tokens) {
		return token.Token{}, false
	}

	t := p.tokens[p.pos]
	p.pos++
	return t, true
}

func (p *baseParser) Accept(ts ...token.Type) {
	token, ok := p.Next()
	if !ok {
		p.Error("no more tokens")
	}

	ttype := token.Type.Value()
	found := false
	for _, t := range ts {
		if t.Value() == ttype {
			found = true
			break
		}
	}
	if !found {
		expectedNames := make([]string, len(ts))
		for i, t := range ts {
			expectedNames[i] = "'" + t.Name() + "'"
		}

		p.Error(fmt.Sprintf("syntax error, expected one of %v but got '%v' at pos %v", expectedNames, token.Type.Name(), token.Pos))
	}
}
