package token

// Stream is an interface providing methods for a usable token stream.
type Stream interface {
	// Tokens returns a channel of Tokens, that can be read from.
	Tokens() <-chan Token
	// Push pushes a token onto this stream. It will be accessible on the
	// channel retrieved with Tokens() eventually.
	Push(Token)
	// Close closes this stream, and also closes the channel retrieved with
	// Tokens().
	Close()
}

// stream is a default implementation of Stream, using a buffered channel with
// buffer size 5 as a queue.
type stream struct {
	ch chan Token
}

// NewStream creates a new, ready to use token stream with a buffer size of 5.
func NewStream() *stream {
	return &stream{
		ch: make(chan Token, 5),
	}
}

// Tokens returns the token channel of this stream.
func (s *stream) Tokens() <-chan Token {
	return s.ch
}

// Push pushes a new token on the channel of this stream.
func (s *stream) Push(t Token) {
	s.ch <- t
}

// Close closes the channel of this stream.
func (s *stream) Close() {
	close(s.ch)
}
