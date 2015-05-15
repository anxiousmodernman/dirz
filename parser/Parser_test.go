package parser

import (
	"fmt"
	"testing"
)

func Test_AddChildTreeItem(t *testing.T) {

	item := treeItem{
		name:        "test",
		indent:      3,
		isDirectory: true,
		children:    nil,
	}

	item2 := treeItem{
		name:        "test2",
		indent:      4,
		isDirectory: true,
		children:    nil,
	}

	item.AddChild(item2)
	if length := len(item.children); length != 1 {
		t.Log("item.children length expected 1, got", length)
		t.Fail()
	}
}

func Test_CreateIndentationStack(t *testing.T) {

	stack := newIndentationStack()

	fmt.Println(stack)
}

func Test_MakeNextId(t *testing.T) {

	parser := Parser{identity: 1}
	parser.nextId()
	if parser.identity != 2 {
		t.Log("Expected 2, got ", parser.identity)
		t.Fail()
	}
}

func Test_Parse3TreeItems(t *testing.T) {
	parser := Parser{identity: 1}
	ParseDirectory(&parser)
}
