package lexer

import (
	// "fmt"
	"github.com/anxiousmodernman/dirz/token"
	"strings"
)

func LexFile(lexer *Lexer) LexFn {
	for {
		_ = "breakpoint"
		if strings.HasPrefix(lexer.InputToEnd(), token.SPACE) {
			// fmt.Println("emitting TOKEN_FILE because of space")
			lexer.Emit(token.TOKEN_FILE)
			ch := lexer.Peek()
			if ch == token.EOF {
				// fmt.Println("token EOF encountered by Lexer")
				lexer.EmitEOF()
				return nil // shutdown signal value for the Run() method
			}
			return LexMeaninglessWhitespace
		}

		if strings.HasPrefix(lexer.InputToEnd(), token.NEWLINE) {
			// fmt.Println("emitting TOKEN_FILE because of newline")
			lexer.Emit(token.TOKEN_FILE)
			return LexNewline
		}

		lexer.Inc()
	}
}
