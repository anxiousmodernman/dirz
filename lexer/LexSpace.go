package lexer

import (
	"github.com/anxiousmodernman/dirz/token"
)

func LexSpace(lexer *Lexer) LexFn {
	for {
		if lexer.IsWhitespace() {
			lexer.Inc()
			lexer.Emit(token.SPACE)
		}

		lexer.Inc()

	}

}
