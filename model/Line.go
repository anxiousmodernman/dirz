package model

type Line struct {
	Indentation  int  // increment to "count" the whitespace and infer a directory hierarchy
	Directory    bool // false implies a file?
	LineNumber   int
	TemplateName string // name of template, if exists
	Name         string // name of directory or file
}
