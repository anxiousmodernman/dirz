package lexer

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/anxiousmodernman/dirz/token"
)

/*
Lexer object contains the state of our parser and provides
a stream for accepting tokens.

Based on work by Rob Pike
http://cuddle.googlecode.com/hg/talk/lex.html#landing-slide
*/
type Lexer struct {
	Name   string
	Input  string
	Tokens chan token.Token
	State  LexFn

	Start int
	Pos   int
	Width int
}

/*
Backup to the beginning of the last read token.
*/
func (this *Lexer) Backup() {
	// this.Debug("Backup()")
	this.Pos -= this.Width
}

/*
Returns a slice of the current input from the current lexer start position
to the current position.
*/
func (this *Lexer) CurrentInput() string {
	return this.Input[this.Start:this.Pos]
}

/*
Decrement the position
*/
func (this *Lexer) Dec() {
	// this.Debug("Dec()")
	this.Pos--
}

/*
Print boring debug info. Pass in a useful name for your context, such as
the name of the calling function.
*/
func (this *Lexer) Debug(context string) {

	// fmt.Println("DEBUG: Context:", context, "Pos =", this.Pos, "; Start =",
	// 	this.Start, "; Width = ", this.Width, "; \n\tInput =", this.InputToEnd())
}

/*
Puts a token onto the token channel. The value of this token is
read from the input based on the current lexer position.
*/
func (this *Lexer) Emit(tokenType token.TokenType) {
	this.Debug("Emit()")
	fmt.Println("emitting this", this.Input[this.Start:this.Pos])
	this.Tokens <- token.Token{Type: tokenType, Value: this.Input[this.Start:this.Pos]}
	this.Start = this.Pos
}

/*
A special EmitEOF fucntion is required, because an index out of bounds error occurs when
accessing a the input string at EOF in regular Emit() function.
*/
func (this *Lexer) EmitEOF() {
	fmt.Println("Emitting EOF")
	this.Tokens <- token.Token{Type: token.TOKEN_EOF, Value: "EOF"}
	this.Shutdown()
}

/*
Returns a token with error information.
*/
func (this *Lexer) Errorf(format string, args ...interface{}) LexFn {
	this.Debug("Errorf()")
	this.Tokens <- token.Token{
		Type:  token.TOKEN_ERROR,
		Value: fmt.Sprintf(format, args...),
	}

	return nil
}

/*
Ignores the current token by setting the lexer's start
position to the current reading position.
*/
func (this *Lexer) Ignore() {
	this.Debug("Ignore")
	this.Start = this.Pos
}

/*
Increment the position
*/
func (this *Lexer) Inc() {
	this.Debug("Inc()")
	this.Pos++
	if this.Pos >= utf8.RuneCountInString(this.Input) {
		this.EmitEOF()
	}
}

/*
Return a slice of the input from the current lexer position
to the end of the input string.
*/
func (this *Lexer) InputToEnd() string {
	if !this.IsEOF() {
		return this.Input[this.Pos:]
	} else {
		return ""
	}
}

/*
Returns the true/false if the lexer is at the end of the
input stream.
*/
func (this *Lexer) IsEOF() bool {
	this.Debug("IsEOF")
	return this.Pos >= len(this.Input)
}

/*
Returns true/false if next character is whitespace
*/
func (this *Lexer) IsWhitespace() bool {
	this.Debug("IsWhitespace()")
	ch, _ := utf8.DecodeRuneInString(this.Input[this.Pos:])
	return unicode.IsSpace(ch)
}

/*
Reads the next rune (character) from the input stream
and advances the lexer position.
*/
func (this *Lexer) Next() rune {
	this.Debug("Next()")
	if this.Pos >= utf8.RuneCountInString(this.Input) {
		this.Width = 0
		return token.EOF
	}

	// get the next run from a string, starting at this.Pos
	result, width := utf8.DecodeRuneInString(this.Input[this.Pos:])

	this.Width = width     // set the byte width
	this.Pos += this.Width // advance this.Pos by this.Width
	return result          // return the rune
}

/*
Return the next token from the channel
*/
func (this *Lexer) NextToken() token.Token {
	for {
		select {
		// either take a token off the channel...
		case token := <-this.Tokens:
			return token
			// ... or CALL the state func and re-set it
		default:
			this.State = this.State(this)
		}
	}

	panic("Lexer.NextToken reached an invalid state!!")
}

/*
Returns the next rune in the stream, then puts the lexer
position back. Basically reads the next rune without consuming
it.
*/
func (this *Lexer) Peek() rune {
	this.Debug("Peek()")
	rune := this.Next()
	this.Backup()
	return rune
}

/*
Starts the lexical analysis and feeding tokens into the
token channel.
*/
func (this *Lexer) Run() {
	for state := LexBegin; state != nil; {
		state = state(this)
	}

	// this.Shutdown()
}

/*
Shuts down the token stream
*/
func (this *Lexer) Shutdown() {
	this.Debug("Shutdown")
	close(this.Tokens)
}

/*
Skips whitespace until we get something meaningful.
*/
func (this *Lexer) SkipWhitespace() {
	this.Debug("SkipWhitespace()")
	for {
		ch := this.Next()

		if !unicode.IsSpace(ch) {
			this.Dec()
			break
		}

		if ch == token.EOF {
			this.Emit(token.TOKEN_EOF)
			break
		}
	}
}
