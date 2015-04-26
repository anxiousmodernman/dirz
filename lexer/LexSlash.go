package lexer

import (
	"github.com/anxiousmodernman/dirz/token"
)

func LexSlash(lexer *Lexer) LexFn {
	lexer.Pos += len(token.SLASH)
	lexer.Emit(token.TOKEN_SLASH) // gotta pass TokenType, not string
	return LexDirectory
}
