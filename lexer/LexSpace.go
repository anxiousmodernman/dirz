package lexer

import (
	"github.com/anxiousmodernman/dirz/token"
)

func LexSpace(lexer *Lexer) LexFn {
	for {
		if lexer.IsWhitespace() {
			lexer.Emit(token.SPACE)
		}

		if strings.HasPrefix(lexer.InputToEnd(), token.SLASH) {
			return LexSlash
		}

		lexer.Inc()

	}

}
