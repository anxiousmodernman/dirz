package parser

import (
	"testing"

	"github.com/anxiousmodernman/dirz/token"
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

func Test_MakeNextId(t *testing.T) {

	parser := Parser{identity: 1}
	parser.nextId()
	if parser.identity != 2 {
		t.Log("Expected 2, got ", parser.identity)
		t.Fail()
	}
}

func Test_Parse3TreeItems(t *testing.T) {

	tkns := make([]token.Token, 10)
	tkns = append(tkns,
		token.Token{Type: token.TOKEN_SLASH, Value: token.SLASH},
		token.Token{Type: token.TOKEN_DIRECTORY_NAME, Value: "dirA"},
		token.Token{Type: token.TOKEN_NEWLINE, Value: token.NEWLINE},
		token.Token{Type: token.TOKEN_SPACE, Value: token.SPACE},
		token.Token{Type: token.TOKEN_SPACE, Value: token.SPACE},
		token.Token{Type: token.TOKEN_SLASH, Value: token.SLASH},
		token.Token{Type: token.TOKEN_DIRECTORY_NAME, Value: "dirB"},
		token.Token{Type: token.TOKEN_NEWLINE, Value: token.NEWLINE},
		token.Token{Type: token.TOKEN_SLASH, Value: token.SLASH},
		token.Token{Type: token.TOKEN_DIRECTORY_NAME, Value: "dirA1"},
		token.Token{Type: token.TOKEN_NEWLINE, Value: token.NEWLINE},
	)

	parser := Parser{
		identity: 1,
		tokens:   tkns,
	}
	ParseDirectory(&parser)
}
