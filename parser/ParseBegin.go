package parser

import (
	"fmt"

	"github.com/anxiousmodernman/dirz/token"
)

func ParseBegin(this *Parser) ParseFn {

	if this.pos > len(this.tokens)-1 {
		return nil
	}

	currentToken := this.tokens[this.pos]
	switch currentToken.Type {

	case token.TOKEN_SPACE:
		fmt.Println("Got TOKEN_SPACE")
		this.currentIndent++
		this.pos++
		return ParseBegin

	case token.TOKEN_SLASH:
		this.previousIndent = this.currentIndent
		fmt.Println("Got TOKEN_SLASH")
		this.pos++
		return ParseDirectory

	case token.TOKEN_FILE:
		this.previousIndent = this.currentIndent
		fmt.Println("Got TOKEN_FILE")
		this.pos++
		return ParseFile

	case token.TOKEN_NEWLINE:
		this.previousIndent = this.currentIndent
		fmt.Println("Got TOKEN_NEWLINE")
		this.pos++
		this.currentLine++
		return ParseBegin

	default:
		fmt.Println("Got invalid input")
	}

	return nil
}
