package parser

import (
	"fmt"
	"io/ioutil"

	"github.com/anxiousmodernman/dirz/lexer"
	"github.com/anxiousmodernman/dirz/token"
)

type parseTree map[int]treeItem

type treeItem struct {
	id           int
	name         string
	indent       int
	isDirectory  bool
	isFile       bool
	templateName string
	parentId     int
	children     []treeItem
}

func (this *treeItem) AddChild(item treeItem) {

	this.children = append(this.children, item)
	fmt.Println("this.children is", this.children)
}

type Parser struct {
	Tree           parseTree
	stack          indentationStack
	currentIndent  int
	previousIndent int
	currentLine    int
	identity       int
	State          ParseFn
	tokens         []token.Token
	pos            int
	// Add a simple slice of map[treeItem.id]indentationLevel
}

func (this *Parser) nextId() {
	this.identity = this.identity + 1
}

func isEOF(theToken token.Token) bool {
	return theToken.Type == token.TOKEN_EOF
}

func (this *Parser) Parse(fileName string) {

	fileContents := readFileToString(fileName)
	var theToken token.Token
	l := lexer.BeginLexing(fileName, fileContents)
	go l.Run()
	var sequence []token.Token

	for {
		theToken = l.NextToken()
		sequence = append(sequence, theToken)
		if isEOF(theToken) {
			break
		}
	}

	this.tokens = sequence
	this.pos = 0
	this.parse()

}

func (this *Parser) parse() {

	for state := ParseBegin; state != nil; {
		state = state(this)
	}

}

func readFileToString(filename string) string {

	fileBytes, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return string(fileBytes)
}
