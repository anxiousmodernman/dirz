package lexer

import (
	// "fmt"
	// "github.com/anxiousmodernman/dirz/errors"
	"strings"

	"github.com/anxiousmodernman/dirz/token"
)

func LexDirectory(lexer *Lexer) LexFn {
	// outside the for loop, because we want to validate the beginning
	if strings.HasPrefix(lexer.InputToEnd(), token.SLASH) {
		panic("Can't have two slashes in a row")
	}

	if strings.HasPrefix(lexer.InputToEnd(), token.SPACE) {
		panic("Gotta have a name for that folder, Tex")
	}

	directoryNameLength := 0

	for {

		if strings.HasPrefix(lexer.InputToEnd(), token.NEWLINE) {
			// fmt.Println("emitting TOKEN_DIRECTORY_NAME")
			lexer.Emit(token.TOKEN_DIRECTORY_NAME)

			return LexNewline
		}

		if lexer.IsEOF() {

			if directoryNameLength == 0 {
				// lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
				lexer.EmitEOF()
			} else {
				lexer.EmitEOF()
			}

			return nil

		}

		lexer.Inc()

		directoryNameLength++
	}
}
