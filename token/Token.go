package token

import "fmt"

type TokenType int

const (
	TOKEN_ERROR TokenType = iota
	TOKEN_EOF

	TOKEN_SLASH
	TOKEN_SPACE
	TOKEN_NAME
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
