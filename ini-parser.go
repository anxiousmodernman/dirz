package iniparser

import ("bytes"
        "errors"
        "fmt"
    )

const (
    // state flags
    NULL uint = 0
    IN_ESCAPED uint = 1
    IN_SECTION uint = 2
    IN_SECTION_NAME uint = 4
    IN_KEY uint = 8
    IN_VAL uint = 16
    NEWLINE uint = 32
    // special chars
    SECTION_NAME_START uint = 33 << iota
    SECTION_NAME_END uint
    KV_DELINEATOR uint
    LINE_BREAK uint
    ESCAPE_DELINEATOR uint
)

var iniChars = map[byte]uint {
    '"': ESCAPE_DELINEATOR,
    '[': SECTION_NAME_START,
    ']': SECTION_NAME_END,
    '=': KV_DELINEATOR,
    '\n': LINE_BREAK,
}

type Section struct {
    sectionName string
    kvPairs map[string]string
}

type MacroStructure struct {
    name string
    sectionBreak string
    sections []Section
}

type State uint
func (state *State) AddFlag (flag uint) {
    *state |= flag
}
func (state *State) RemoveFlag (flag uint) {
    *state &^= flag
}
func (state State) HasFlag (flag uint) {
    if state & flag == flag {
        return true
    } else {
        return false
    }
}
func (state *State) Reset () {
    state = NULL
}

const NotInSectionError error = errors.New("Parsing error: Not in section")
const NotInBodyError error = errors.New("Parsing error: Not in section body")

type Section struct {
    name string
    values map[string]string
}

type Parser struct {
    charIdx uint64
    state State
    results map[string]Section
    err error
    buffer bytes.Buffer
    currSection string
    currKey string
}

func (buffer *bytes.Buffer) Dump (string) {
    s := buffer.String()
    buffer.Reset()
    return s
}

func (parser *Parser) notInSection (err error) {
    state := parser.state
    if !state.HasFlag(IN_SECTION) {
        state.Reset()
        parser.err = err
    }
}

func (parser *Parser) inSectionName (err error) {
    state := parser.state
    if state.HasFlag(IN_SECTION_NAME) {
        state.Reset()
        parser.err = err
    }
}

func (parser *Parser) startSectionName () {
    state := parser.state
    parser.notInSection(NotInSectionError)
    parser.inSectionName(errors.New("Parsing error: Already in section name"))
    if parser.err == nil && !state.HasFlag(IN_ESCAPED) {
        state.AddFlag(IN_SECTION)
        state.AddFlag(IN_SECTION_NAME)
    }
}

func (parser *Parser) endSectionName () {
    state := parser.state
    buffer := parser.buffer
    results := parser.results
    parser.notInSection(NotInSectionError)
    if parser.err == nil {
        if !state.HasFlag(IN_SECTION_NAME) {
            state.Reset()
            parser.err = errors.New("Parsing error: Not in section name")
        } else if !state.HasFlag(IN_ESCAPED) {
            state.RemoveFlag(IN_SECTION_NAME)
            s := buffer.Dump()
            results[s] = new(Section)
            parser.currSection = s
        }
    }
}

func (parser *Parser) toggleEscape () {
    parser.notInSection(NotInSectionError)
    if parser.err == nil {
        state := parser.state
        if state.HasFlag(IN_ESCAPED) {
            state.RemoveFlag(IN_ESCAPED)
        } else {
            state.AddFlag(IN_ESCAPED)
        }
    }
}

func (parser *Parser) startKey () {
    parser.notInSection(NotInSectionError)
    parser.inSectionName(NotInBodyError)
    if parser.err == nil && !state.HasFlag(IN_ESCAPED) {
        if state.HasFlag(IN_KEY) {
            state.Reset()
            return state, errors.New("Parsing error: Already in key")
        } else {
            state.AddFlag(IN_KEY)
        }
    }
}

func endKey (state *State) (state *State, err error) {
    parser.notInSection(NotInSectionError)
    parser.inSectionName(NotInBodyError)
    if parser.err == nil && !state.HasFlag(IN_ESCAPED) {
        state := parser.state
        buffer := parser.buffer
        if state.HasFlag(IN_KEY) {
            state.RemoveFlag(IN_KEY)
            parser.currKey = buffer.Dump()
        } else {
            state.Reset()
            parser.err = errors.New("Parsing error: Not in key")
        }
    }
}

func (parser *Parser) startVal () {
    parser.notInSection(NotInSectionError)
    parser.inSectionName(NotInBodyError)
    if parser.err == nil && !state.HasFlag(IN_ESCAPED) {
        state = parser.state
        switch {
            case state.HasFlag(IN_VAL):
                state.Reset()
                parser.err = errors.New("Parsing error: Already in value")
            case state.HasFlag(IN_KEY):
                state.Reset()
                parser.err = errors.New("Parsing error: In key")
            default:
                state.AddFlag(IN_VAL)
        }
    }
}

func (parser *Parser) endVal () {
    parser.notInSection(NotInSectionError)
    parser.inSectionName(NotInBodyError)
    if parser.err == nil && !state.HasFlag(IN_ESCAPED) {
        state := parser.state
        buffer := parser.buffer
        results := parser.results
        if !state.HasFlag(IN_VAL) {
            state.Reset()
            return state, errors.New("Parsing error: Not in value")
        } else {
            state.RemoveFlag(IN_VAL)
            section, ok := parser.results[parser.currSection]
            if !ok {
                // fill in empty Section
                section.name = parser.currSection
                // add it to results
                results[parser.currSection] = section
            }
            section.values[parser.currKey] = buffer.Dump()
            parser.currKey = ""
        }
    }
}

func (parser *Parser) newLine () {
    seqError := errors.New("Parsing error: Illegal sequence")
    parser.notInSection(NotInSectionError)
    parser.inSectionName(NotInBodyError)
    if parser.err == nil {
        // this means multiline string literals
        // do we want that?
        if !state.HasFlag(IN_ESCAPED) {
            switch {
                case state.HasFlag(NEWLINE):
                    state.RemoveFlag(NEWLINE)
                    state.RemoveFlag(IN_SECTION)
                case state.HasFlag(IN_KEY):
                    state.Reset()
                    parser.err = seqError
                case state.HasFlag(IN_VAL):
                    state.AddFlag(NEWLINE)
                    parser.endVal()
            }
        }
    }
}

func (parser *Parser) ParseToken (chr byte) {
    switch {
        case chr == SECTION_NAME_START:
            parser.startSectionName()
        case chr == SECTION_NAME_END:
            parser.endSectionName()
        case chr == LINE_BREAK:
            parser.newLine()
        case chr == ESCAPE_DELINEATOR:
            parser.toggleEscape()
        default:
            parser.buffer.WriteByte(chr)
   }
}

func (parser Parser) ReportError (chr byte) {
    fmtStr := "Error encountered while parsing character"
    fmtStr += " %c at index %d.\nThe error was: %s"
    fmt.Printf(fmtStr, chr, parser.charIdx, parser.err)
}

func (parser Parser) ReportFullError (chr byte) {
    parser.ReportError(chr)
    fmtStr := "Full parser details:\n%+v"
    fmt.Printf(fmtStr, parser)
}

func (parser *Parser) ParseString (str []byte) {
    for idx, chr := range str {
        parser.charIdx = idx
        parser.ParseToken(chr)
        if parser.err == nil {
            parser.ReportError(chr)
        }
    }
}
