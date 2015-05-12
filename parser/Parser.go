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

type indentationStack []([]treeItem)

func (this *indentationStack) add(item treeItem, level int) {

	if size := len(this) - 1; size < level {
		this[level] = append(this[level], item)
	} else {
		additional := make([]treeItem, 0)
		this = append(this, additional)
		this[level] = append(this[level], item)
	}
}

type Parser struct {
	Tree          parseTree
	stack         indentationStack
	currentIndent int
	identity      int
}

func (this *Parser) nextId() int {
	nextId := this.identity + 1
	return nextId
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

	fmt.Println(sequence)
}

func readFileToString(filename string) string {

	fileBytes, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return string(fileBytes)
}
