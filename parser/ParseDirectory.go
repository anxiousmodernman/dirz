package parser

import (
	"fmt"
	"os"

	"github.com/anxiousmodernman/dirz/token"
)

func ParseDirectory(this *Parser) ParseFn {

	directory := this.tokens[this.pos]
	if directory.Type != token.TOKEN_DIRECTORY_NAME {
		fmt.Println("Parser error on line", this.currentLine, "- Expected directory name")
		os.Exit(1)
	}

	this.nextId()
	item := treeItem{
		id:           this.identity,
		name:         directory.Value,
		indent:       this.currentIndent,
		isDirectory:  true,
		isFile:       false,
		templateName: "",
		parentId:     -1, // determine if it has a parent
	}
	if this.currentIndent < this.previousIndent {
		// add to the previous indentation stack
	}
	this.Tree[item.id] = item
	fmt.Println("printing parseTree", this.Tree)

	// find out if there is a template or newline

	for {
		this.pos++
		tkn := this.tokens[this.pos]
		if tkn.Type == token.TOKEN_EOF {
			return nil
		}

		if tkn.Type == token.TOKEN_NEWLINE {
			return ParseBegin
		}

	}

	return nil
}
