package lexer

import (
	"fmt"
	"github.com/anxiousmodernman/dirz/token"
	"strings"
)

func LexDirectory(lexer *Lexer) LexFn {

	for {
		if strings.HasPrefix(lexer.InputToEnd(), token.SLASH) {
			panic("Can't have two slashes in a row")
		}

		if strings.HasPrefix(lexer.InputToEnd(), token.SPACE) {
			panic("Gotta have a name for that folder, Tex")
		}

		if strings.HasPrefix(lexer.InputToEnd(), token.NEWLINE) {
			fmt.Println("emitting TOKEN_DIRECTORY_NAME")
			lexer.Emit(token.TOKEN_DIRECTORY_NAME)
			return LexBegin
		}

		lexer.Inc()
	}
}
