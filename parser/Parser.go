package parser

import (
	"fmt"
	"io/ioutil"

	"github.com/anxiousmodernman/dirz/lexer"
	"github.com/anxiousmodernman/dirz/model"
	"github.com/anxiousmodernman/dirz/token"
)

type Parser struct {
	Tree               parseTree
	currentIndentation int
}

func isEOF(theToken token.Token) bool {
	return theToken.Type == token.TOKEN_EOF
}

func (this *Parser) Parse(fileName string) parseTree {

	fileContents := readFileToString(fileName)
	var theToken token.Token
	output := parseTree

	fmt.Println("Starting lexer and parser for file", fileName, "...")

	l := lexer.BeginLexing(fileName, fileContents)
	go l.Run()

	var sequence []token.Token

	for {
		theToken = l.NextToken()

		sequence = append(sequence, theToken)

		if isEOF(theToken) {
			fmt.Println("Parser encountered EOF")
			break
		}

	}

	fmt.Println(sequence)

	fmt.Println("Parser has been shutdown")
	return output
}

func readFileToString(filename string) string {

	fileBytes, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return string(fileBytes)
}
