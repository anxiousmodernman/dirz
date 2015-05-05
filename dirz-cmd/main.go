package main

import (
	// "fmt"
	dirz "github.com/anxiousmodernman/dirz/parser"
)

func main() {

	var filename = "../testfiles/touch-these-files.dirz"
	dirz.ParseFile(filename)
	// const threeSpaces = "   "
	// for index, runeValue := range threeSpaces {
	// 	fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	// }
}
