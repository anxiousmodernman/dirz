package parser

import (
	"fmt"
	"os"

	"github.com/anxiousmodernman/dirz/token"
)

func ParseDirectory(this *Parser) ParseFn {

	_ = "breakpoint"
	directory := this.tokens[this.pos]
	if directory.Type != token.TOKEN_DIRECTORY_NAME {
		fmt.Println("Parser error on line", this.currentLine, "- Expected directory name")
		os.Exit(1)
	}

	// add treeItem to parseTree

	//

	return nil
}
