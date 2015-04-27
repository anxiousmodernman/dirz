package dirz

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "rakerparser"
    "rakerlexer"
)
type Flags struct {
    dryRun bool
    bufSize int
    inputFilePath string
    rootDir string
}

func makeFlags () (flags Flags) {
    cwd, _ := os.Getwd()
    flags = Flags{}
    flag.BoolVar(&flags.dryRun, "d", false,
                  "Dry run--print directory structure instead of creating it")
    flag.IntVar(&flags.bufSize, "b", 3, "Buffer size")
    flag.StringVar(&flags.inputFilePath, "i", "", "Path to input file")
    flag.StringVar(&flags.rootDir, "r", cwd,
                   "Starting directory (defaults to current directory)")
    return flags
}

func main() {
    // get CLI args
    flags := makeFlags()
    flag.Parse()
    inputFilePath := flag.Arg(0)
    if inputFilePath == "" {
        inputFilePath = flags.inputFilePath
    }
    rootDir := flags.rootDir

    var errString = fmt.Sprintf("Error reading file %s", flags.inputFilePath)

    // initialize lexer & parser
    tokens := make(chan rakerlexer.Token)
    lexer := rakerlexer.NewLexer(tokens)
    parser := rakerparser.NewParser(rootDir, tokens)

    // read input file
    inputFile, err := os.Open(inputFilePath)
    if err != nil {
        print(errString)
    } else {
        rdr := bufio.NewReader(inputFile)
        go lexer.ParseString(rdr)
        go parser.StartParsing()
    }

    //block until we're done parsing
    <-parser.Done

    // dryRun -> just printing
    if flags.dryRun {
        println(parser.DirectoryTree.String())
    } else {
        parser.DirectoryTree.MakeDirectory(rootDir)
        print("Done")
    }
}
