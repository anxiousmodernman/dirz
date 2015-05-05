package lexer

import (
	"github.com/anxiousmodernman/dirz/token"
	"testing"
)

func countTokenType(tokens []token.Token, tokenType token.TokenType) (count int) {
	for _, tkn := range tokens {
		if tkn.Type == tokenType {
			count += 1
		}
	}
	return
}

func Test_LexSingleLine(t *testing.T) {

	input := `  /`
	lxr := BeginLexing("testLexer", input)
	go lxr.Run()

	var tokens []token.Token

	// Normally this for loop would be internal to the parser, but we can fake it for testing
	for {
		thing := lxr.NextToken()
		tokens = append(tokens, thing)
		if thing.Type == token.TOKEN_EOF {
			break
		}
	}

	if len(tokens) != 4 {
		// Why len 4? Answer: two spaces, a slash, and EOF token
		t.Log("Expected tokens slice to be length 4, got ", len(tokens))
		t.Fail()
	}
	if tokens[0].Type != token.TOKEN_SPACE {
		t.Logf("Expected TOKEN_SPACE, got %s", tokens[0].Type)
		t.Fail()
	}
	if tokens[3].Type != token.TOKEN_EOF {
		t.Logf("Expected TOKEN_EOF, got %s", tokens[3].Type)
		t.Fail()
	}
}

func Test_Lex4LinesWithDirectoryOnFirstLine(t *testing.T) {

	input :=
		`/first
  /second
  file1
  file3
`
	lxr := BeginLexing("testLexer", input)
	go lxr.Run()

	var tokens []token.Token

	// Normally this for loop would be internal to the parser, but we can fake it for testing
	for {
		thing := lxr.NextToken()
		tokens = append(tokens, thing)
		if thing.Type == token.TOKEN_EOF {
			break
		}
	}

	if newLines := countTokenType(tokens, token.TOKEN_NEWLINE); newLines != 4 {
		t.Log("Expected 4 newlines, got ", newLines)
		t.Fail()
	}
	if slashes := countTokenType(tokens, token.TOKEN_SLASH); slashes != 2 {
		t.Log("Expected 2 slashes, got ", slashes)
		t.Fail()
	}
	if directoryNames := countTokenType(tokens, token.TOKEN_DIRECTORY_NAME); directoryNames != 2 {
		t.Log("Expected 2 names, got ", directoryNames)
		t.Fail()
	}
	if fileNames := countTokenType(tokens, token.TOKEN_FILE); fileNames != 2 {
		t.Log("Expected 2 file names got ", fileNames)
		t.Fail()
	}
	if eof := countTokenType(tokens, token.TOKEN_EOF); eof != 1 {
		t.Log("Expected 1 EOF, got ", eof)
	}
}

func Test_LexWithNewlineAsFirstToken(t *testing.T) {

	input :=
		`
/first
  /second
  file1
  file3
`
	lxr := BeginLexing("testLexer", input)
	go lxr.Run()

	var tokens []token.Token

	// Normally this for loop would be internal to the parser, but we can fake it for testing
	for {
		thing := lxr.NextToken()
		tokens = append(tokens, thing)
		if thing.Type == token.TOKEN_EOF {
			break
		}
	}

	if newLines := countTokenType(tokens, token.TOKEN_NEWLINE); newLines != 5 {
		t.Log("Expected 5 newlines, got ", newLines)
		t.Fail()
	}
	if slashes := countTokenType(tokens, token.TOKEN_SLASH); slashes != 2 {
		t.Log("Expected 2 slashes, got ", slashes)
		t.Fail()
	}
	if directoryNames := countTokenType(tokens, token.TOKEN_DIRECTORY_NAME); directoryNames != 2 {
		t.Log("Expected 2 names, got ", directoryNames)
		t.Fail()
	}
	if fileNames := countTokenType(tokens, token.TOKEN_FILE); fileNames != 2 {
		t.Log("Expected 2 file names got ", fileNames)
		t.Fail()
	}
	if eof := countTokenType(tokens, token.TOKEN_EOF); eof != 1 {
		t.Log("Expected 1 EOF, got ", eof)
	}
}
