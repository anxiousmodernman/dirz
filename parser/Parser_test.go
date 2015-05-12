package parser

import "testing"

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
