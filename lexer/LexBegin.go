package lexer

import (
	"fmt"
	"strings"

	"github.com/anxiousmodernman/dirz/token"
)

func LexBegin(lexer *Lexer) LexFn {

	fmt.Println("LexBegin called")

	if strings.HasPrefix(lexer.InputToEnd(), token.SPACE) {
		return LexSpace
	} else if strings.HasPrefix(lexer.InputToEnd(), token.SLASH) {
		fmt.Println("LexBegin: SLASH matched")

		return LexSlash
	} else if len(lexer.InputToEnd()) == 0 {
		panic("empty file")
	} else {
		panic("LexBegin matched nothing.")
	}

}
