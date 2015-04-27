package rakerparser

import (
    "fmt"
    "os"
    "path/filepath"
)

func (directory Directory) MakeDirectory (curPath string) {
    newDirPath := filepath.Join(curPath, directory.name)
    os.MkdirAll(newDirPath, os.ModeDir)
    for _, dir := range directory.children {
        dir.MakeDirectory(newDirPath)
    }
}

func (directory Directory) toString (s []byte, tabCount int) (out string) {
    dirName := fmt.Sprintf("%s/\n", directory.name)
    line := []byte{}
    for i := 0; i < tabCount; i++ {
        copy(line, "\t")
    }
    copy(line, dirName)
    copy(s, line)
    tc := tabCount + 1
    for _, dir := range directory.children {
        dir.toString(s, tc)
    }
    out = string(s)
    return out
}

func (directory Directory) String () string {
    return directory.toString([]byte{}, 0)
}
