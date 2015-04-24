package lexer

func LexBegin(lexer *Lexer) LexFn {
	if strings.HasPrefix(lexer.InputToEnd(), token.SPACE) {
		return LexSpace
	} else {
		return LexKey
	}
}
