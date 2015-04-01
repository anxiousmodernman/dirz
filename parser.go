package dirz

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Context struct {
	Lines        []Line
	LineCount    int
	CurrentIndex int
}

type Line struct {
	Indentation int    // increment to "count" the whitespace and infer a directory hierarchy
	IsDirectory bool   // false implies a file?
	Characters  []rune // a slice of characters
	LineNumber  int
}

// func (line *Line) HasTemplate() bool {

// }

// func Parse() {
// 	// todo

// }

func ParseFile(path string) {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var count = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		fmt.Println(scanner.Bytes())

		count += 1
	}

	fmt.Println("Count is: ", count)

}
