package parser

import (
	"fmt"
	"github.com/anxiousmodernman/dirz/token"
)

func ParseBegin(this *Parser) ParseFn {

	currentToken := this.tokens[this.pos]

	switch currentToken.Type {
	case token.TOKEN_SPACE:
		fmt.Println("Got TOKEN_SPACE")
		this.pos++
		return ParseBegin
	case token.TOKEN_SLASH:
		this.pos++
		fmt.Println("Got TOKEN_SLASH")
		return ParseDirectory
	case token.TOKEN_FILE:
		fmt.Println("Got TOKEN_FILE")
		this.pos++
		return ParseFile
	case token.TOKEN_NEWLINE:
		fmt.Println("Got TOKEN_NEWLINE")
		this.pos++
		return ParseBegin
	default:
		fmt.Println("Got invalid input")
	}
}