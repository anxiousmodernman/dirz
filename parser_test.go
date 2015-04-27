package dirz

import (
	"github.com/anxiousmodernman/dirz/parser"
	"testing"
)

func TestParseSingleFile(t *testing.T) {
	var filename = "./testfiles/touch-these-files.dirz"
	_ = "breakpoint"
	ctx := parser.Parse(filename)

	if ctx.LineCount != 4 {
		t.Fail()
		t.Log("Line count should be 4, got", ctx.LineCount)
	}

}
