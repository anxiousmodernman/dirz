package lexer

import (
	"github.com/anxiousmodernman/dirz/token"
	"strings"
)

func LexFile(lexer *Lexer) LexFn {
	for {
		if strings.HasPrefix(lexer.InputToEnd(), token.SPACE) {
			lexer.Emit(token.TOKEN_FILE)
			ch := lexer.Peek()
			if ch == token.EOF {
				lexer.EmitEOF()
				return nil // shutdown signal value for the Run() method
			}
			return LexMeaninglessWhitespace
		}

		if strings.HasPrefix(lexer.InputToEnd(), token.NEWLINE) {
			lexer.Emit(token.TOKEN_FILE)
			return LexNewline
		}

		lexer.Inc()
	}
}
