package lexer

import (
	// "fmt"
	"strings"

	"github.com/anxiousmodernman/dirz/token"
)

func LexBegin(lexer *Lexer) LexFn {

	if strings.HasPrefix(lexer.InputToEnd(), token.SPACE) {
		return LexSpace
	} else if strings.HasPrefix(lexer.InputToEnd(), token.SLASH) {
		return LexSlash
	} else if strings.HasPrefix(lexer.InputToEnd(), token.NEWLINE) {
		return LexNewline
	} else if lexer.IsEOF() {
		lexer.EmitEOF()
		return nil // breaks out of Run() loop
	} else {
		panic("LexBegin matched nothing.")
	}

}
