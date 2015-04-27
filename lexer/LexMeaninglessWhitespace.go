package lexer

import (
	"fmt"
	"github.com/anxiousmodernman/dirz/token"
	"strings"
)

func LexMeaninglessWhitespace(lexer *Lexer) LexFn {

	lexer.SkipWhitespace()

	if strings.HasPrefix(lexer.InputToEnd(), token.NEWLINE) {
		fmt.Println("NEWLINE MATCHED inside LexMeaninglessWhitespace")
		return LexNewline
	} else {
		// Eventually we will support specifying templates after the file or directory, but not yet
		panic("non-space characters not allowed after file or directory name")
	}

}
