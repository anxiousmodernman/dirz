package dirz

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"unicode/utf8"
)

const dir = '/'

func ParseFile(path string) Context {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines = 0
	var ctx = Context{nil, lines, 0}
	//var parsedLines []Line

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines += 1
		b := scanner.Bytes()
		r, _ := utf8.DecodeRune(b)
		if isWhitespace(r) {
			fmt.Println("some whitespace here", "bytes:", b, "rune:", r, "bytes len", len(b))

		}
		// parsed := ParseLine(scanner.Text(), lines)
		// parsedLines = append(parsedLines, parsed)
	}

	ctx.LineCount = lines
	fmt.Println(lines, " lines scanned")

	return ctx
}

func ParseLine(line string, lineNumber int) Line {

	var parsed = Line{Indentation: 0, Directory: false, Characters: nil, LineNumber: lineNumber}
	var nonWhitespaceScanned = false
	var chars []rune
	var leadingWhitespace = 0

	for _, runeValue := range line {
		chars = append(chars, runeValue)

		if ws := isWhitespace(runeValue); ws == true && !nonWhitespaceScanned {
			leadingWhitespace += 1
		} else {
			nonWhitespaceScanned = true
		}
	}

	parsed.Characters = chars
	parsed.Indentation = leadingWhitespace
	parsed.Directory, _ = hasDirectory(parsed)
	fmt.Println(parsed)

	return parsed

}

func hasDirectory(line Line) (bool, error) {
	//var dirIndex int
	var dirSeen = false
	for _, char := range line.Characters {
		if char == dir {
			if dirSeen == true {
				log.Fatal("Parse error. Multiple \"/\" directory identifiers found on line ", line.LineNumber)
				return true, errors.New("Parse error")
			}

		}
	}
	return false, nil
}

func isWhitespace(r rune) bool {

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
