package main

import (
	"fmt"

	"github.com/TimSatke/parser/lexer"
)

// Character classes
const (
	CCHexChar      = lexer.StringCharacterClass("ABCDEFabcdef1234567890")
	CCNumberStart  = lexer.StringCharacterClass("1234567890-")
	CCNumber       = lexer.StringCharacterClass("1234567890")
	CCSignNeg      = lexer.StringCharacterClass("-")
	CCSign         = lexer.StringCharacterClass("+-")
	CCExponent     = lexer.StringCharacterClass("Ee")
	CCFraction     = lexer.StringCharacterClass(".")
	CCAfterEscape  = lexer.StringCharacterClass("\"\\/bfnrtu")
	CCSingleQuote  = lexer.StringCharacterClass("\"")
	CCBracketOpen  = lexer.StringCharacterClass("{")
	CCBracketClose = lexer.StringCharacterClass("}")
	CCBraceOpen    = lexer.StringCharacterClass("[")
	CCBraceClose   = lexer.StringCharacterClass("]")
	CCColon        = lexer.StringCharacterClass(":")
	CCComma        = lexer.StringCharacterClass(",")
	CCWhitespace   = lexer.StringCharacterClass("\u0020\u000D\u000A\u0009")
)

func lexJSON(l lexer.Lexer) lexer.State {
	return lexToken
}

func lexToken(l lexer.Lexer) lexer.State {
	switch r := l.Peek(); r {
	case '{':
		return lexBracketOpen
	case '}':
		return lexBracketClose
	case '[':
		return lexBraceOpen
	case ']':
		return lexBraceClose
	case '"':
		return lexString
	case ':':
		return lexColon
	case ',':
		return lexComma
	case 't':
		return lexTrue
	case 'f':
		return lexFalse
	case 'n':
		return lexNull
	default:
		// handle all cases that cannot be expressed directly in a switch
		if CCWhitespace.Matches(r) {
			return lexWhitespace
		}
		if CCNumberStart.Matches(r) {
			return lexNumber
		}
	}
	return unexpectedToken
}

func lexNumber(l lexer.Lexer) lexer.State {
	l.Accept(CCSignNeg) // optional

	if !l.AcceptMultiple(CCNumber) {
		return tokenMismatch(l, CCNumber)
	}

	if l.Accept(CCFraction) {
		if !l.AcceptMultiple(CCNumber) {
			return tokenMismatch(l, CCNumber)
		}
	}

	if l.Accept(CCExponent) {
		l.Accept(CCSign)
		if !l.AcceptMultiple(CCNumber) {
			return tokenMismatch(l, CCNumber)
		}
	}

	l.Emit(TokenNumber)

	return lexToken
}

func lexTrue(l lexer.Lexer) lexer.State {
	expectation := "true"

	for _, c := range expectation {
		next := l.Next()
		if c != next {
			l.Backup()
			return unexpectedToken
		}
	}
	l.Emit(TokenTrue)

	return lexToken
}

func lexFalse(l lexer.Lexer) lexer.State {
	expectation := "false"

	for _, c := range expectation {
		next := l.Next()
		if c != next {
			l.Backup()
			return unexpectedToken
		}
	}
	l.Emit(TokenFalse)

	return lexToken
}

func lexNull(l lexer.Lexer) lexer.State {
	expectation := "null"

	for _, c := range expectation {
		next := l.Next()
		if c != next {
			l.Backup()
			return unexpectedToken
		}
	}
	l.Emit(TokenNull)

	return lexToken
}

func lexString(l lexer.Lexer) lexer.State {
	if !l.Accept(CCSingleQuote) {
		return tokenMismatch(l, CCSingleQuote)
	}

	escaped := false

loop:
	for {
		r := l.Next()

		if escaped && !CCAfterEscape.Matches(r) {
			l.Backup()
			return tokenMismatch(l, CCAfterEscape)
		}

		switch r {
		case '"':
			if !escaped {
				break loop
			}
			escaped = !escaped
		case '\\':
			escaped = !escaped
		case 'u':
			if escaped {
				for i := 0; i < 4; i++ {
					if !l.Accept(CCHexChar) {
						return tokenMismatch(l, CCHexChar)
					}
				}
				escaped = false
				break
			}
			fallthrough
		default:
			escaped = false
		}
	}

	l.Emit(TokenString)

	return lexToken
}

func lexColon(l lexer.Lexer) lexer.State {
	if !l.Accept(CCColon) {
		return tokenMismatch(l, CCColon)
	}
	l.Emit(TokenColon)
	return lexToken
}

func lexComma(l lexer.Lexer) lexer.State {
	if !l.Accept(CCComma) {
		return tokenMismatch(l, CCComma)
	}
	l.Emit(TokenComma)
	return lexToken
}

func lexBracketOpen(l lexer.Lexer) lexer.State {
	if !l.Accept(CCBracketOpen) {
		return tokenMismatch(l, CCBracketOpen)
	}
	l.Emit(TokenBracketOpen)
	return lexToken
}

func lexBracketClose(l lexer.Lexer) lexer.State {
	if !l.Accept(CCBracketClose) {
		return tokenMismatch(l, CCBracketClose)
	}
	l.Emit(TokenBracketClose)
	return lexToken
}

func lexBraceOpen(l lexer.Lexer) lexer.State {
	if !l.Accept(CCBraceOpen) {
		return tokenMismatch(l, CCBraceOpen)
	}
	l.Emit(TokenBraceOpen)
	return lexToken
}

func lexBraceClose(l lexer.Lexer) lexer.State {
	if !l.Accept(CCBraceClose) {
		return tokenMismatch(l, CCBraceClose)
	}
	l.Emit(TokenBraceClose)
	return lexToken
}

func lexWhitespace(l lexer.Lexer) lexer.State {
	l.AcceptMultiple(CCWhitespace)
	l.Emit(TokenWhitespace)

	return lexToken
}

func tokenMismatch(l lexer.Lexer, expected lexer.CharacterClass) lexer.State {
	return errorf(l, "Unexpected token, expected one of [%s], got '%s'", expected.String(), string(l.Peek()))
}

func unexpectedToken(l lexer.Lexer) lexer.State {
	return errorf(l, "Unexpected token '%s'", string(l.Peek()))
}

func errorf(l lexer.Lexer, msg string, args ...interface{}) lexer.State {
	l.EmitError(TokenError, fmt.Sprintf(msg, args...))
	return nil
}
