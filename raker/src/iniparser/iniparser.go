package iniparser

import "inilexer"

type DirectoryMap map[string]Directory

type Directory struct {
    name string
    parentDirectory *Directory
    children DirectoryMap
}

func (directory *Directory) AddSubDirectory (subDir *Directory) {
    *subDir.parentDirectory = *directory
    directory.children[subDir.name] = *subDir
}

func NewDirectory(dirName string) (newDir Directory) {
    newDir = Directory{dirName, nil, make(DirectoryMap)}
    return newDir
}

type Parser struct {
    directoryTree *Directory
    topDirs []*Directory
    tokens chan inilexer.Token
    maxTabs uint
    tabCount uint
}

func NewParser (rootDirName string) (newParser Parser) {
    rootDir := NewDirectory(rootDirName)
    return Parser{&rootDir, []*Directory{&rootDir},
                  make(chan inilexer.Token), 0, 0}
}

func (parser *Parser) PopDir () {
    topDirs := parser.topDirs
    if len(topDirs) > 0 {
        // get the last dir in topDirs
        poppedDir := topDirs[len(topDirs) - 1]
        // truncate it
        topDirs = topDirs[:len(topDirs) - 1]
        /* get the *new* last dir, i.e. the parent
           dir of poppedDir */
        topDir := topDirs[len(topDirs) - 1]
        topDir.AddSubDirectory(poppedDir)
    }
}

func (parser *Parser) StopParsing () {
    for (len(parser.topDirs) > 0) {
        parser.PopDir()
    }
}

func (parser *Parser) ParseTab () {
    parser.tabCount++
}

func (parser *Parser) ParseName (dirName string) {
    newDir := NewDirectory(dirName)
    tabDiff := parser.maxTabs - parser.tabCount
    for i := uint(0); i <= tabDiff; i++ {
        parser.PopDir()
    }
    parser.topDirs = append(parser.topDirs, &newDir)
    parser.maxTabs = parser.tabCount
    parser.tabCount = 0
}

func (parser *Parser) ParseToken(token inilexer.Token) {
    switch {
        case token.TokenType == inilexer.TAB:
                parser.ParseTab()
        case token.TokenType == inilexer.DIRECTORY_NAME:
                parser.ParseName(token.Content)
        case token.TokenType == inilexer.EOF:
                parser.StopParsing()
        }
}

func (parser *Parser) StartParsing () {
    Parse:
        for {
            token := <-parser.tokens
            if (token.TokenType == inilexer.ERROR) {
                break Parse
            } else {
                parser.ParseToken(token)
            }
        }
}
