package rakerparser

import (
    "rakerlexer"
    "path/filepath"
)

type DirectorySlice []Directory

type Directory struct {
    name string
    parentDirectory *Directory
    children DirectorySlice
}

func (directory *Directory) AddSubDirectory (subDir *Directory) {
    *subDir.parentDirectory = *directory
    directory.children = append(directory.children, *subDir)
}

func (directory Directory) GetPath () (path string) {
    parents := []string{}
    for {
        parent := directory.parentDirectory
        if parent == nil {
            break
        } else {
            parents = append(parents, parent.name)
        }
    }
    return filepath.Join(parents...)
}

func NewDirectory(dirName string) (newDir Directory) {
    newDir = Directory{dirName, nil, DirectorySlice{}}
    return newDir
}

type Parser struct {
    DirectoryTree *Directory
    topDirs []*Directory
    tokens chan rakerlexer.Token
    Done chan bool
    maxTabs uint
    tabCount uint
}

func NewParser (rootDirName string, tokens chan rakerlexer.Token) (newParser Parser) {
    rootDir := NewDirectory(rootDirName)
    return Parser{&rootDir, []*Directory{&rootDir}, tokens, make(chan bool), 0, 0}
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
    parser.Done <- true
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

func (parser *Parser) ParseToken(token rakerlexer.Token) {
    switch {
        case token.TokenType == rakerlexer.TAB:
                parser.ParseTab()
        case token.TokenType == rakerlexer.DIRECTORY_NAME:
                parser.ParseName(token.Content)
        case token.TokenType == rakerlexer.EOF:
                parser.StopParsing()
        }
}

func (parser *Parser) StartParsing () {
    Parse:
        for {
            token := <-parser.tokens
            if (token.TokenType == rakerlexer.ERROR) {
                break Parse
            } else {
                parser.ParseToken(token)
            }
        }
}
