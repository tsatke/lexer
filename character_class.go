package lexer

import "strings"

// CharacterClass is an interface providing methods for matching runes.
//
// Implementations are
//
//  lexer.StringCharacterClass
//  lexer.NotStringCharacterClass
//
// Both of these can be used to define constant character classes. See their
// documentation for more information.
type CharacterClass interface {
	// Matches returns whether the given rune is matched by this character
	// class.
	Matches(rune) bool
	String() string
}

// StringCharacterClass is an implementation of lexer.CharacterClass, which
// matches runes that are contained in the string used to define the character
// class.
//
//  const WhitespaceNoLinefeed = lexer.StringCharacterClass(" \t") // will match all runes that are either ' ' or '\t'
type StringCharacterClass string

// Matches returns true if the given rune is contained inside the definition of
// this character class.
func (s StringCharacterClass) Matches(r rune) bool {
	return strings.IndexRune(string(s), r) >= 0
}

func (s StringCharacterClass) String() string { return string(s) }

// NotStringCharacterClass is an implementation of lexer.CharacterClass, which
// matches runes that are NOT contained in the string used to define the
// character class.
//
//  const WhitespaceNoLinefeed = lexer.StringCharacterClass(" \t") // will match all runes that are neither ' ' nor '\t'
type NotStringCharacterClass string

// Matches returns true if the given rune is NOT contained inside the definition of
// this character class.
func (s NotStringCharacterClass) Matches(r rune) bool {
	return strings.IndexRune(string(s), r) < 0
}

func (s NotStringCharacterClass) String() string { return string(s) }
