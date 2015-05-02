package lexer

import (
	// "fmt"
	"strings"

	"github.com/anxiousmodernman/dirz/token"
)

func LexBegin(lexer *Lexer) LexFn {

	// fmt.Println("LexBegin called")

	if strings.HasPrefix(lexer.InputToEnd(), token.SPACE) {
		// fmt.Println("LexBegin: SPACE matched Start", lexer.Start, "Pos", lexer.Pos, "Width", lexer.Width)
		return LexSpace
	} else if strings.HasPrefix(lexer.InputToEnd(), token.SLASH) {
		// fmt.Println("LexBegin: SLASH matched", lexer.Start, "Pos", lexer.Pos, "Width", lexer.Width)
		return LexSlash
	} else if lexer.IsEOF() {
		// fmt.Println("LexBegin: EOF matched", lexer.Start, "Pos", lexer.Pos, "Width", lexer.Width)
		lexer.EmitEOF()
		return nil // breaks out of Run() loop
	} else {
		panic("LexBegin matched nothing.")
	}

}
