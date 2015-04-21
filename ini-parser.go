package iniparser

import ("errors"
        "bytes")

const parserState (
    NULL uint = 0
    IN_ESCAPED = uint 1
    IN_SECTION uint = 2
    IN_SECTION_NAME uint = 4
    IN_KEY uint = 8
    IN_VAL uint = 16
    NEWLINE uint = 32
)

const ParseChars (
    SECTION_NAME_START uint = iota
    SECTION_NAME_END uint
    KV_DELINEATOR uint
    LINE_BREAK uint
    ESCAPE_DELINEATOR uint
)

var iniChars = map[byte]uint {
    '"': ESCAPE_DELINEATOR
    '[': SECTION_NAME_START
    ']': SECTION_NAME_END
    '=': KV_DELINEATOR
    '\n': LINE_BREAK
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

func inSection (state *uint) bool {
    if state.HasFlag(IN_SECTION) {
        return true
    } else {
        state.Reset()
        return false
    }
}

func inSectionName (state *uint) bool {
    if state.HasFlag(IN_SECTION_NAME) {
        state.Reset()
        return true
    } else {
        return false
    }
}

const NotInSectionError error = errors.New("Parsing error: Not in section")
const NotInBodyError error = errors.New("Parsing error: Not in section body")

type Section struct {
    name string
    values map[string][string]
}

type Parser struct {
    text bytes.Buffer
    charIdx uint64
    state State
    results map[bytes.Buffer]Section
    err error
    buffer bytes.Buffer
}

func (parser *Parser) startSection () {
    state = parser.state
    switch {
        case state.HasFlag(IN_ESCAPED) {
            return state, nil
        }
        case state.HasFlag(IN_SECTION) {
            state.Reset()
            return state, errors.New("Parsing error: Already in section")
        }
        case inSectionName(state) {
            return state, NotInBodyError
        }
        default {
            state.AddFlag(IN_SECTION)
            return state, nil
        }
    }
}

func (parser *Parser) endSection () {
    state = parser.state
    switch {
        case state.HasFlag(IN_ESCAPED) {
            return state, nil
        }
        case !inSection(state) {
            return state, NotInSectionError
        }
        case inSectionName(state) {
            return state, NotInBodyError
        }
        default {
            state.RemoveFlag(IN_SECTION)
            return state, nil
        }
    }
}

func (parser *Parser) startSectionName () {
    state = parser.State
    switch {
        case state.HasFlag(IN_ESCAPED) {
            return state, nil
        }
        case !inSection(state) {
            return state, NotInSectionError
        }
        case inSectionName(state) {
            return state, errors.New("Parsing error: Already in section name")
        }
        default {
            state.AddFlag(IN_SECTION_NAME)
            return state, nil
        }
    }
}

func (parser *Parser) endSectionName () {
    state = parser.State
    buffer = parser.buffer
    results = parser.results
    switch {
        case state.HasFlag(IN_ESCAPED) {
            // ???
        }
        case !inSection(state) {
            return state, NotInSectionError
        }
        case !state.HasFlag(IN_SECTION_NAME) {
            state.Reset()
            return state, errors.New("Parsing error: Not in section name")
        }
        default {
            state.RemoveFlag(IN_SECTION_NAME)
            s := buffer.String()
            results[s] = new(Section)
            buffer.Reset()
            return state, nil
        }
    }
}

func toggleEscape (state *State) (state *State, err error) {
    switch {
        case !inSection(state) {
            return state, NotInSectionError
        }
        case state.HasFlag(IN_ESCAPED) {
            state.RemoveFlag(IN_ESCAPED)
            return state, nil
        }
        default {
            state.AddFlag(IN_ESCAPED)
            return state, nil
        }
    }
}

func startKey (state *State) (state *State, err error) {
    switch {
        case state.HasFlag(IN_ESCAPED) {
            return state, nil
        }
        case inSectionName(state) {
            return state, NotInBodyError
        }
        case !inSection(state) {
            return state, errors.New("Parsing error: Not in section")
        }
        case state.HasFlag(IN_KEY) {
            state.Reset()
            return state, errors.New("Parsing error: Already in key")
        }
        default {
            state.AddFlag(IN_KEY)
            return state, nil
        }
    }
}

func endKey (state *State) (state *State, err error) {
    switch {
        case state.HasFlag(IN_ESCAPED) {
            return state, nil
        }
        case inSectionName(state) {
            return state, NotInBodyError
        }
        case !inSection(state) {
            return state, NotInSectionError
        }
        case state.HasFlag(IN_KEY) {
            state.RemoveFlag(IN_KEY)
            return state, nil
        }
        default {
            state.Reset()
            return state, errors.New("Parsing error: Not in key")
        }
    }
}

func startVal (state *State) (state *State, err error) {
    switch {
        case state.HasFlag(IN_ESCAPED) {
            return state, nil
        }
        case !inSection(state) {
            return state, NotInSectionError
        }
        case inSectionName(state) {
            return state, NotInBodyError
        }
        case state.HasFlag(IN_VAL) {
            state.Reset()
            return state, errors.New("Parsing error: Already in value")
        }
        case state.HasFlag(IN_KEY) {
            state.Reset()
            return state, errors.New("Parsing error: In key")
        }
        default {
            state.AddFlag(IN_VAL)
            return state, nil
        }
    }
}

func endVal (state *State) (state *State, err error) {
    switch {
        case state.HasFlag(IN_ESCAPED) {
            return state, nil
        }
        case !inSection(state) {
            return state, NotInSectionError
        }
        case state.HasFlag(IN_SECTION_NAME) {
            state.Reset()
            return state, NotInBodyError
        }
        case !state.HasFlag(IN_VAL) {
            state.Reset()
            return state, errors.New("Parsing error: Not in value")
        }
        // don't need to check for state.HasFlag(IN_KEY)
        default {
            state.RemoveFlag(IN_VAL)
            return state, nil
        }
    }
}

func newLine (state *State) (state *State, err error) {
    seqError := errors.New("Parsing error: Illegal sequence")
    switch {
        case !inSection(state) {
            return state, NotInSectionError
        }
        // this means multiline string literals
        // do we want that?
        case state.HasFlag(IN_ESCAPED) {
            return state, nil
        }
        case state.HasFlag(NEWLINE) {
            state.RemoveFlag(NEWLINE)
            state.RemoveFlag(IN_SECTION)
            return state, nil
        }
        case state.HasFlag(IN_SECTION_NAME) {
            state.Reset()
            return state, seqError
        }
        /* the next case validates section contents formatted as:
           [sectionName]
           keyName=valName
           
           if we want lists in addition to/instead of maps, this
           should be 
           case state.HasFlag(IN_KEY) {
               state.RemoveFlag(IN_KEY)
               return state, seqError
           }
        */
        case state.HasFlag(IN_KEY) {
            state.Reset()
            return state, seqError
        }
        case state.HasFlag(IN_VAL) {
            state.RemoveFlag(IN_VAL)
            state.AddFlag(NEWLINE)
            return state, nil
        }
    }
}


var dispatchMap = map[uint]stateChange {
    SECTION_NAME_START: startSectionName
    SECTION_NAME_END: endSectionName
    ESCAPE_DELINEATOR: toggleEscape
    LINE_BREAK: newLine
}

func (parser *Parser) parseToken (chr byte) {
    state = *parser.state
    results = *parser.results
    f, ok := dispatchMap[chr]
    if ok {
        _, err := f(state)
        if err != nil {
            parser.err = err
        }
    } else {
        switch {
            case state.HasFlag(IN_SECTION_NAME) {
               secName 



