package rakerlexer

import (
    "bytes"
    "io"
    "unicode"
)

type TokenType uint8

const (
    // state flags
    STOPPED uint8 = 0
    IN_ERROR uint8 = 1
    IN_ESCAPED uint8 = 2
    IN_NAME uint8 = 4
    // special chars
    LINE_BREAK rune = '\n'
    ESCAPE_DELINEATOR rune = '"'
    SLASH rune = '/'
    REPLACEMENT_CHAR rune = unicode.ReplacementChar
    // token types
    ERROR TokenType = 0
    TAB TokenType = 1
    DIRECTORY_NAME TokenType = 2
    EOF TokenType = 4
)

type Token struct {
    TokenType TokenType
    Content string
}

type State uint8
func (state *State) AddFlag (flag uint8) {
    s := uint8(*state)
    s |= flag
    *state = State(s)
}
func (state *State) RemoveFlag (flag uint8) {
    s := uint8(*state)
    s &^= flag
    *state = State(s)
}
func (state State) HasFlag (flag uint8) bool {
    s := uint8(state)
    if s & flag == flag {
        return true
    } else {
        return false
    }
}

type Lexer struct {
    charIdx int
    whitespaceCount uint8
    state State
    output chan Token
    buffer *bytes.Buffer
}

func (lexer *Lexer) EmitError () {
    lexer.output <- Token{TokenType: ERROR, Content: "Invalid sequence"}
    lexer.state.AddFlag(IN_ERROR)
    close(lexer.output)
}

func (lexer *Lexer) Emit (token Token) {
    lexer.output <- token
}

func (lexer *Lexer) WhiteSpace () {
    if !lexer.state.HasFlag(IN_ESCAPED) {
        if lexer.whitespaceCount == 3 {
            lexer.Emit(Token{TokenType: TAB, Content: ""})
            lexer.whitespaceCount = 0
        } else {
            lexer.whitespaceCount++
        }
    }
}

func (lexer *Lexer) Slash () {
    if !lexer.state.HasFlag(IN_ESCAPED) {
        if lexer.state.HasFlag(IN_NAME) {
            lexer.EmitError()
        } else {
            lexer.state.AddFlag(IN_NAME)
        }
    }
}

func (lexer *Lexer) NewLine () {
    if lexer.state.HasFlag(IN_NAME) {
        s := lexer.buffer.String()
        lexer.buffer.Reset()
        token := Token{TokenType: DIRECTORY_NAME, Content: s}
        lexer.Emit(token)
    } else {
        // no multiline escaped strings
        lexer.EmitError()
    }
}

func (lexer *Lexer) ToggleEscape () {
    if lexer.state.HasFlag(IN_ESCAPED) {
        lexer.state.RemoveFlag(IN_ESCAPED)
    } else {
        lexer.state.AddFlag(IN_ESCAPED)
    }
}

func (lexer *Lexer) ParseToken (chr rune) {
    switch {
        case chr == SLASH:
            lexer.Slash()
        case chr == LINE_BREAK:
            lexer.NewLine()
        case chr == ESCAPE_DELINEATOR:
            lexer.ToggleEscape()
        // returned when io.RuneReader can't read rune
        case chr == REPLACEMENT_CHAR:
            lexer.EmitError()
        default:
            lexer.buffer.WriteRune(chr)
    }
}

func (lexer *Lexer) ParseString (rdr io.RuneReader) {
    lexer.state.RemoveFlag(STOPPED)
    idx := 0
    Parse:
        for {
            if (lexer.state.HasFlag(STOPPED) || lexer.state.HasFlag(IN_ERROR)) {
                break Parse
            } else {
                chr, _, _ := rdr.ReadRune()
                lexer.charIdx = idx
                lexer.ParseToken(chr)
                idx++
            }
        }
    lexer.Emit(Token{EOF, ""})
}

func NewLexer(tokens chan Token) (l Lexer) {
    l = Lexer{
            charIdx: 0,
            state: State(STOPPED),
            output: tokens,
            buffer: new(bytes.Buffer),
        }
    return l
}
