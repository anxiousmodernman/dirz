package parser

import (
	"fmt"
	"io/ioutil"

	"github.com/anxiousmodernman/dirz/lexer"
	"github.com/anxiousmodernman/dirz/model"
	"github.com/anxiousmodernman/dirz/token"
)

func isEOF(theToken token.Token) bool {
	return theToken.Type == token.TOKEN_EOF
}

func Parse(fileName string) model.DirzContext {

	fileContents := readFileToString(fileName)
	var theToken token.Token
	// var tokenValue string

	output := model.DirzContext{LineCount: 0}

	fmt.Println("Starting lexer and parser for file", fileName, "...")

	// in LexerFactory.go; returns *Lexer
	l := lexer.BeginLexing(fileName, fileContents)
	// start Run() loop in goroutine
	// first method call is LexBegin, returns a StateFn,
	// repeatedly calls StateFn that return StateFn like this:
	//   state = state(this)
	// the state functions will do their thing and asynchronously put tokens on a chan of Token
	go l.Run()

	var sequence []token.Token

	for {
		theToken = l.NextToken()

		sequence = append(sequence, theToken)

		if isEOF(theToken) {
			fmt.Println("Parser encountered EOF")
			// if hasSection == true {
			// 	fmt.Println("Adding section '", section.Name, "' to output...")
			// 	output.Sections = append(output.Sections, section)
			// }

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
