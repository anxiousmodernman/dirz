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

func ParseFile(path string) Context {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines = 0
	var ctx = Context{nil, lines, 0}
	var parsedLines []Line

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines += 1
		parsed := ParseLine(scanner.Text(), lines)
		parsedLines = append(parsedLines, parsed)
	}

	ctx.LineCount = lines
	fmt.Println(lines, " lines scanned")

	return ctx
}

func ParseLine(line string, lineNumber int) Line {

	var parsed = Line{Indentation: 0, IsDirectory: false, Characters: nil, LineNumber: lineNumber}
	var nonWhitespaceScanned = false
	var chars []rune
	var leadingWhitespace = 0

	for _, runeValue := range line {
		chars = append(chars, runeValue)

		if ws := IsWhitespace(runeValue); ws == true && !nonWhitespaceScanned {
			leadingWhitespace += 1
		} else {
			nonWhitespaceScanned = true
		}
	}

	parsed.Characters = chars
	parsed.Indentation = leadingWhitespace
	fmt.Println(parsed)

	return parsed

}

func IsWhitespace(r rune) bool {

	result := false

	switch r {
	case
		'\u0009', // horizontal tab
		'\u000A', // line feed
		// '\u000B', // vertical tab
		// '\u000C', // form feed
		// '\u000D', // carriage return
		'\u0020', // space
		'\u0085', // next line
		'\u00A0', // no-break space
		'\u1680', // ogham space mark
		'\u180E', // mongolian vowel separator
		'\u2000', // en quad
		'\u2001', // em quad
		'\u2002', // en space
		'\u2003', // em space
		'\u2004', // three-per-em space
		'\u2005', // four-per-em space
		'\u2006', // six-per-em space
		'\u2007', // figure space
		'\u2008', // punctuation space
		'\u2009', // thin space
		'\u200A', // hair space
		'\u2028', // line separator
		'\u2029', // paragraph separator
		'\u202F', // narrow no-break space
		'\u205F', // medium mathematical space
		'\u3000': // ideographic space
		result = true
	default:
		result = false
	}
	return result
}
