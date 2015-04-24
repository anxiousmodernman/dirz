package dirz

type Line struct {
	Indentation int    // increment to "count" the whitespace and infer a directory hierarchy
	Directory   bool   // false implies a file?
	Characters  []rune // a slice of characters
	LineNumber  int
}
