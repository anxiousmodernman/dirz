package lexer

import (
	"github.com/anxiousmodernman/dirz/token"
	"strings"
)

func LexMeaninglessWhitespace(lexer *Lexer) LexFn {

	lexer.SkipWhitespace()

	if strings.HasPrefix(lexer.InputToEnd(), token.NEWLINE) {
		return LexNewline
	} else {
		return LexFile
	}

}
