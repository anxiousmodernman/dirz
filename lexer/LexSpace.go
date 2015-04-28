package lexer

import (
	"fmt"
	"github.com/anxiousmodernman/dirz/token"
	"strings"
)

func LexSpace(lexer *Lexer) LexFn {
	fmt.Println("LexSpace called")
	var nextFn LexFn
	for {
		if lexer.IsWhitespace() {
			fmt.Println("LexSpace: IsWhitespace matched. Start", lexer.Start, "Pos", lexer.Pos, "Width", lexer.Width)
			lexer.Pos += len(token.SPACE)
			lexer.Emit(token.TOKEN_SPACE)

		} else if strings.HasPrefix(lexer.InputToEnd(), token.SLASH) {
			fmt.Println("LexSpace: SLASH matched")
			nextFn = LexSlash
			break
		} else {
			fmt.Println("LexSpace: Returning LexFile")
			nextFn = LexFile
			break
		}

	}
	fmt.Println("now calling", nextFn)
	return nextFn

}
