package lexer

import "strings"

type CharacterClass interface {
	Matches(rune) bool
	String() string
}

type StringCharacterClass string

func (s StringCharacterClass) Matches(r rune) bool {
	return strings.IndexRune(string(s), r) >= 0
}

func (s StringCharacterClass) String() string { return string(s) }

type NotStringCharacterClass string

func (s NotStringCharacterClass) Matches(r rune) bool {
	return strings.IndexRune(string(s), r) < 0
}

func (s NotStringCharacterClass) String() string { return string(s) }
