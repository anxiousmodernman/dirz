package lexer

import (
	"fmt"
	"github.com/anxiousmodernman/dirz/token"
	"strings"
)

func LexDirectory(lexer *Lexer) LexFn {
	// outside the for loop, because we want to validate the beginning
	if strings.HasPrefix(lexer.InputToEnd(), token.SLASH) {
		panic("Can't have two slashes in a row")
	}

	if strings.HasPrefix(lexer.InputToEnd(), token.SPACE) {
		panic("Gotta have a name for that folder, Tex")
	}

	for {

		if strings.HasPrefix(lexer.InputToEnd(), token.NEWLINE) {
			fmt.Println("emitting TOKEN_DIRECTORY_NAME")
			lexer.Emit(token.TOKEN_DIRECTORY_NAME)

			return LexNewline
		}

		lexer.Inc()
	}
}
