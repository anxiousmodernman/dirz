package lexer

import (
	"fmt"
	"github.com/anxiousmodernman/dirz/token"
	"strings"
)

func LexSpace(lexer *Lexer) LexFn {
	fmt.Println("LexSpace called")
	for {
		if lexer.IsWhitespace() {
			fmt.Println("LexSpace: IsWhitespace matched")
			lexer.Pos += len(token.SPACE)
			lexer.Emit(token.TOKEN_SPACE)

		} else if strings.HasPrefix(lexer.InputToEnd(), token.SLASH) {
			fmt.Println("LexSpace: SLASH matched")
			return LexSlash
		} else {
			return LexFile
		}

	}

}
