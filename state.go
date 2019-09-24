package lexer

// State is a recursive definition of lexer states, that are executed
// non-recursively.
//
// See the following example. The goal is, to lex strings that match exactly
// ABC. This should be tokenized into three tokens, TokenA, TokenB and TokenC.
// The following example shows, how to define states to achieve this (without
// error handling, just to show the sequence).
//
//	const (
//		TokenA MyTokenType = iota
//		TokenB
//		TokenC
//	)
//	const (
//		CCA = lexer.StringCharacterClass("A")
//		CCB = lexer.StringCharacterClass("B")
//		CCC = lexer.StringCharacterClass("C")
//	)
//	func lexABCString(l lexer.Lexer) lexer.State {
//		return lexA
//	}
//	func lexA(l lexer.Lexer) lexer.State {
//		l.Accept(CCA)
//		l.Emit(TokenA)
//		return lexB
//	}
//	func lexB(l lexer.Lexer) lexer.State {
//		l.Accept(CCB)
//		l.Emit(TokenB)
//		return lexC
//	}
//	func lexC(l lexer.Lexer) lexer.State {
//		l.Accept(CCC)
//		l.Emit(TokenC)
//		return nil
//	}
//
// The lexer will start with lexABCString (assuming that this is the start State
// that you passed when creating the lexer), which will be executed. The lexer
// passed in is the lexer you are working with. lexABCString does nothing with
// the lexer, and returns lexA as next state. The lexer will execute lexA next.
// The lexer passed in is the same lexer as for lexABCString. lexA accepts an
// "A", emits a TokenA and returns lexB. The lexer will now execute lexB, which
// does almost the same as lexA. lexB the returns lexC, which will cause the
// lexer to execute lexC next. lexC returns nil as next state, which tells the
// lexer that the state machine is done, and it will stop execution and close
// the token stream.
type State func(Lexer) State
