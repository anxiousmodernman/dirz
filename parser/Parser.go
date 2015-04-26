package parser

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/anxiousmodernman/dirz/lexer"
	"github.com/anxiousmodernman/dirz/model"
	"github.com/anxiousmodernman/dirz/token"
)

func isEOF(theToken token.Token) bool {
	return theToken.Type == token.TOKEN_EOF
}

func Parse(fileName string) model.DirzContext {

	fileContents := readFileToString(fileName)
	var token token.Token
	var tokenValue string

	output := model.DirzContext{LineCount: 0}

	fmt.Println("Starting lexer and parser for file", fileName, "...")

	// in LexerFactory.go; returns *Lexer
	l := lexer.BeginLexing(fileName, fileContents)
	// start Run() loop in goroutine
	// first method call is LexBegin, returns a StateFn,
	// repeatedly calls StateFn that return StateFn like this:
	//   state = state(this)
	// the state functions will do their thing and asynchronously put tokens on a chan of Token
	// go l.Run()

	for {
		theToken = l.NextToken()

		if theToken.Type != token.TOKEN_VALUE {
			tokenValue = strings.TrimSpace(token.Value)
		} else {
			tokenValue = token.Value
		}

		if isEOF(theToken) {
			// if hasSection == true {
			// 	fmt.Println("Adding section '", section.Name, "' to output...")
			// 	output.Sections = append(output.Sections, section)
			// }

			break
		}

		// switch token.Type {
		// case token.TOKEN_SECTION:
		// 	/*
		// 	 * Reset tracking variables
		// 	 */
		// 	if hasSection == true {
		// 		fmt.Println("Adding section '", section.Name, "' to output...")
		// 		output.Sections = append(output.Sections, section)
		// 	}

		// 	key = ""
		// 	hasSection = true

		// 	section.Name = tokenValue
		// 	section.KeyValuePairs = make([]ini.IniKeyValue, 0)

		// 	fmt.Println("Section", section.Name, "started...")

		// case token.TOKEN_KEY:
		// 	key = tokenValue
		// 	fmt.Println("Key:", key)

		// case token.TOKEN_VALUE:
		// 	fmt.Println("Value:", tokenValue)

		// 	section.KeyValuePairs = append(section.KeyValuePairs, ini.IniKeyValue{Key: key, Value: tokenValue})
		// 	key = ""
		// }
	}

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
