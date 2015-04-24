package inilexer

import ("bytes"
        "errors"
        "fmt"
    )

const (
    // state flags
    NULL rune = 0
    IN_ESCAPED rune = 1
    IN_SECTION rune = 2
    IN_SECTION_NAME rune = 4
    IN_KEY rune = 8
    IN_VAL rune = 16
    NEWLINE rune = 32
    // special chars
    SECTION_NAME_START rune = '['
    SECTION_NAME_END rune = ']'
    KV_DELINEATOR rune = '='
    LINE_BREAK rune = '\n'
    ESCAPE_DELINEATOR rune = '"'
)

type Section struct {
    sectionName string
    kvPairs map[string]string
}

type MacroStructure struct {
    name string
    sectionBreak string
    sections []Section
}

type State rune
func (state *State) AddFlag (flag rune) {
    s := rune(*state)
    s |= flag
    *state = State(s)
}
func (state *State) RemoveFlag (flag rune) {
    s := rune(*state)
    s &^= flag
    *state = State(s)
}
func (state State) HasFlag (flag rune) bool {
    s := rune(state)
    if s & flag == flag {
        return true
    } else {
        return false
    }
}
func (state *State) Reset () {
    *state = State(NULL)
}

var NotInSectionError error = errors.New("Parsing error: Not in section")
var NotInBodyError error = errors.New("Parsing error: Not in section body")

func dump (buffer bytes.Buffer) (s string) {
    s = buffer.String()
    buffer.Reset()
    return s
}

type Lexer struct {
    charIdx int
    state State
    results map[string]Section
    err error
    buffer bytes.Buffer
    currSection string
    currKey string
}
func (lexer *Lexer) notInSection (err error) {
    state := &lexer.state
    if !state.HasFlag(IN_SECTION) {
        state.Reset()
        lexer.err = err
    }
}

func (lexer *Lexer) inSectionName (err error) {
    state := &lexer.state
    if state.HasFlag(IN_SECTION_NAME) {
        state.Reset()
        lexer.err = err
    }
}

func (lexer *Lexer) startSectionName () {
    state := &lexer.state
    lexer.notInSection(NotInSectionError)
    lexer.inSectionName(errors.New("Parsing error: Already in section name"))
    if lexer.err == nil && !state.HasFlag(IN_ESCAPED) {
        state.AddFlag(IN_SECTION)
        state.AddFlag(IN_SECTION_NAME)
    }
}

func (lexer *Lexer) endSectionName () {
    state := &lexer.state
    buffer := lexer.buffer
    results := lexer.results
    lexer.notInSection(NotInSectionError)
    if lexer.err == nil {
        if !state.HasFlag(IN_SECTION_NAME) {
            state.Reset()
            lexer.err = errors.New("Parsing error: Not in section name")
        } else if !state.HasFlag(IN_ESCAPED) {
            state.RemoveFlag(IN_SECTION_NAME)
            s := dump(buffer)
            results[s] = *new(Section)
            lexer.currSection = s
        }
    }
}

func (lexer *Lexer) toggleEscape () {
    lexer.notInSection(NotInSectionError)
    if lexer.err == nil {
        state := &lexer.state
        if state.HasFlag(IN_ESCAPED) {
            state.RemoveFlag(IN_ESCAPED)
        } else {
            state.AddFlag(IN_ESCAPED)
        }
    }
}

func (lexer *Lexer) startKey () {
    lexer.notInSection(NotInSectionError)
    lexer.inSectionName(NotInBodyError)
    state := &lexer.state
    if lexer.err == nil && !state.HasFlag(IN_ESCAPED) {
        if state.HasFlag(IN_KEY) {
            state.Reset()
            lexer.err = errors.New("Parsing error: Already in key")
        } else {
            state.AddFlag(IN_KEY)
        }
    }
}

func (lexer *Lexer) endKey () {
    lexer.notInSection(NotInSectionError)
    lexer.inSectionName(NotInBodyError)
    state := &lexer.state
    if lexer.err == nil && !state.HasFlag(IN_ESCAPED) {
        state := &lexer.state
        buffer := lexer.buffer
        if state.HasFlag(IN_KEY) {
            state.RemoveFlag(IN_KEY)
            lexer.currKey = dump(buffer)
        } else {
            state.Reset()
            lexer.err = errors.New("Parsing error: Not in key")
        }
    }
}

func (lexer *Lexer) startVal () {
    lexer.notInSection(NotInSectionError)
    lexer.inSectionName(NotInBodyError)
    state := &lexer.state
    if lexer.err == nil && !state.HasFlag(IN_ESCAPED) {
        switch {
            case state.HasFlag(IN_VAL):
                state.Reset()
                lexer.err = errors.New("Parsing error: Already in value")
            case state.HasFlag(IN_KEY):
                state.Reset()
                lexer.err = errors.New("Parsing error: In key")
            default:
                state.AddFlag(IN_VAL)
        }
    }
}

func (lexer *Lexer) endVal () {
    lexer.notInSection(NotInSectionError)
    lexer.inSectionName(NotInBodyError)
    state := &lexer.state
    if lexer.err == nil && !state.HasFlag(IN_ESCAPED) {
        buffer := lexer.buffer
        results := lexer.results
        if !state.HasFlag(IN_VAL) {
            state.Reset()
            lexer.err = errors.New("Parsing error: Not in value")
        } else {
            state.RemoveFlag(IN_VAL)
            section, ok := lexer.results[lexer.currSection]
            if !ok {
                // fill in empty Section
                section.sectionName = lexer.currSection
                // add it to results
                results[lexer.currSection] = section
            }
            section.kvPairs[lexer.currKey] = dump(buffer)
            lexer.currKey = ""
        }
    }
}

func (lexer *Lexer) newLine () {
    seqError := errors.New("Parsing error: Illegal sequence")
    lexer.notInSection(NotInSectionError)
    lexer.inSectionName(NotInBodyError)
    if lexer.err == nil {
        state := &lexer.state
        // this means multiline string literals
        // do we want that?
        if !state.HasFlag(IN_ESCAPED) {
            switch {
                case state.HasFlag(NEWLINE):
                    state.RemoveFlag(NEWLINE)
                    state.RemoveFlag(IN_SECTION)
                case state.HasFlag(IN_KEY):
                    state.Reset()
                    lexer.err = seqError
                case state.HasFlag(IN_VAL):
                    state.AddFlag(NEWLINE)
                    lexer.endVal()
            }
        }
    }
}

func (lexer *Lexer) ParseToken (chr rune) {
    switch {
        case chr == SECTION_NAME_START:
            lexer.startSectionName()
        case chr == SECTION_NAME_END:
            lexer.endSectionName()
        case chr == LINE_BREAK:
            lexer.newLine()
        case chr == ESCAPE_DELINEATOR:
            lexer.toggleEscape()
        default:
            lexer.buffer.WriteRune(chr)
   }
}

func (lexer Lexer) ReportError (chr rune) {
    fmtStr := "Error encountered while parsing character"
    fmtStr += " %c at index %d.\nThe error was: %s"
    fmt.Printf(fmtStr, chr, lexer.charIdx, lexer.err)
}

func (lexer Lexer) ReportFullError (chr rune) {
    lexer.ReportError(chr)
    fmtStr := "Full lexer details:\n%+v"
    fmt.Printf(fmtStr, lexer)
}

func (lexer *Lexer) ParseString (str []rune) {
    for idx, chr := range str {
        lexer.charIdx = idx
        lexer.ParseToken(chr)
        if lexer.err != nil {
            lexer.ReportError(chr)
            break
        }
    }
}
