package lexer

import (
	"github.com/anxiousmodernman/dirz/token"
	"strings"
)

func LexSpace(lexer *Lexer) LexFn {
	for {
		if lexer.IsWhitespace() {
			lexer.Emit(token.TOKEN_SPACE)
		}

		if strings.HasPrefix(lexer.InputToEnd(), token.SLASH) {
			return LexSlash
		}

		lexer.Inc()

	}

}
