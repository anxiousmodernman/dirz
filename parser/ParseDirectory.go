package parser

import (
	"fmt"
	"github.com/anxiousmodernman/dirz/token"
	"os"
)

func ParseDirectory(this *Parser) ParseFn {
	directory := this.tokens[this.pos]
	if directory.Type != token.TOKEN_DIRECTORY_NAME {
		fmt.Println("Parser error on line", this.currentLine, "- Expected directory name")
		os.Exit(1)
	}

	// add treeItem to parseTree

	//
}

func parseDirectoryWithIndent()
