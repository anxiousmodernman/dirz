package lexer

import (
	"fmt"
	"github.com/anxiousmodernman/dirz/token"
	"testing"
)

func TestLexSingleToken(t *testing.T) {

	input := `  /`
	lxr := BeginLexing("testLexer", input)
	_ = "breakpoint"
	go lxr.Run()

	var tokens []token.Token

	for {

		thing := lxr.NextToken()
		tokens = append(tokens, thing)

		fmt.Println("Token: ", thing.Value)
		if thing.Type == token.TOKEN_EOF {
			break
		}

	}

	t.Log(tokens)

}
