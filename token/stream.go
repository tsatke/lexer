package token

type Stream interface {
	Tokens() <-chan Token
	Push(Token)
	Close()
}

type stream struct {
	ch chan Token
}

func NewStream() *stream {
	return &stream{
		ch: make(chan Token, 5),
	}
}

func (s *stream) Tokens() <-chan Token {
	return s.ch
}

func (s *stream) Push(t Token) {
	s.ch <- t
}

func (s *stream) Close() {
	close(s.ch)
}
