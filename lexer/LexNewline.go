package lexer

import (
	// "fmt"
	"github.com/anxiousmodernman/dirz/token"
)

func LexNewline(lexer *Lexer) LexFn {
	// fmt.Println("Calling LexNewline")
	lexer.Pos += len(token.NEWLINE)
	lexer.Emit(token.TOKEN_NEWLINE)
	return LexBegin
}
