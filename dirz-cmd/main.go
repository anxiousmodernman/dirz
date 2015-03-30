package main

import "fmt"

func main() {
	const threeSpaces = "   "
	for index, runeValue := range threeSpaces {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
}
