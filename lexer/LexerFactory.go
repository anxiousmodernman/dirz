package lexer

import (
	"github.com/anxiousmodernman/dirz/token"
)

/*
Start a new lexer with a given input string. This returns the
instance of the lexer and a channel of tokens. Reading this stream
is the way to parse a given input and perform processing.
*/
func BeginLexing(name string, input string) *Lexer {
	lxr := &Lexer{
		Name:   name,
		Input:  input,
		State:  LexBegin,
		Tokens: make(chan token.Token),
	}

	return lxr
}
