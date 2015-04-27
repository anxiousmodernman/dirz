package lexer

import (
	"fmt"
	"strings"

	"github.com/anxiousmodernman/dirz/token"
)

func LexBegin(lexer *Lexer) LexFn {

	fmt.Println("LexBegin called")

	if strings.HasPrefix(lexer.InputToEnd(), token.SPACE) {
		fmt.Println("LexBegin: SPACE matched")
		return LexSpace
	} else if strings.HasPrefix(lexer.InputToEnd(), token.SLASH) {
		fmt.Println("LexBegin: SLASH matched")
		return LexSlash
	} else if lexer.IsEOF() {
		fmt.Println("LexBegin: EOF matched")
		lexer.Emit(token.TOKEN_EOF)
		return nil // breaks out of Run() loop
	} else {
		panic("LexBegin matched nothing.")
	}

}
