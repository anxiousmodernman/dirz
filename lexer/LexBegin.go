package lexer

import (
	"strings"

	"github.com/anxiousmodernman/dirz/token"
)

func LexBegin(lexer *Lexer) LexFn {

	if strings.HasPrefix(lexer.InputToEnd(), token.SPACE) {
		return LexSpace
	} else {
		return LexBegin
	}

}
