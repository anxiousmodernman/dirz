package parser

import (
	"fmt"

	"github.com/anxiousmodernman/dirz/token"
)

func ParseBegin(this *Parser) ParseFn {

	this.previousIndent = this.currentIndent

	if this.pos > len(this.tokens)-1 {
		return nil
	}

	currentToken := this.tokens[this.pos]
	//_ = "breakpoint"
	switch currentToken.Type {

	case token.TOKEN_SPACE:
		fmt.Println("Got TOKEN_SPACE")
		this.pos++
		this.cur
		return ParseBegin

	case token.TOKEN_SLASH:
		fmt.Println("Got TOKEN_SLASH")
		this.pos++
		return ParseDirectory

	case token.TOKEN_FILE:
		fmt.Println("Got TOKEN_FILE")
		this.pos++
		return ParseFile

	case token.TOKEN_NEWLINE:
		fmt.Println("Got TOKEN_NEWLINE")
		this.pos++
		this.currentLine++
		return ParseBegin

	default:
		fmt.Println("Got invalid input")
	}

	return nil // todo
}
