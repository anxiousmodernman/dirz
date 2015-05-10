package parser

type parseTree map[treeItem]treeItem

type treeItem struct {
	name         string
	indent       int
	isDirectory  bool
	isFile       bool
	templateName string
}
