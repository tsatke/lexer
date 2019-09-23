package parser

import (
	"github.com/TimSatke/parser/parser/ast"
	"github.com/TimSatke/parser/token"
)

type Rule func(Parser) Rule

type Parser interface {
	Parse() error

	Tree() ast.Tree

	Error(string)

	HasMore() bool
	Peek() (token.Token, bool)
	Next() (token.Token, bool)
	Accept(...token.Type)
}
