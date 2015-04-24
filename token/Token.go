package token

import "fmt"

type TokenType int

const (
	TOKEN_ERROR TokenType = iota
	TOKEN_EOF

	TOKEN_LEFT_BRACKET
	TOKEN_RIGHT_BRACKET
	TOKEN_EQUAL_SIGN
	TOKEN_NEWLINE

	TOKEN_SECTION
	TOKEN_KEY
	TOKEN_VALUE
)

type Token struct {
	Type  TokenType
	Value string
}

func (this Token) String() string {
	switch this.Type {
	case TOKEN_EOF:
		return "EOF"

	case TOKEN_ERROR:
		return this.Value
	}

	return fmt.Sprintf("%q", this.Value)
}
