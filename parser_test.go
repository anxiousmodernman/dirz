package dirz

import "testing"

func TestParseSingleFile(t *testing.T) {
	var filename = "./testfiles/touch-these-files.dirz"

	// file should be all like
	// /level0
	//   file1
	//   filt2
	//   file3

	var ctx Context
	ctx = ParseFile(filename)

	if ctx.LineCount != 4 {
		t.Fail()
		t.Log("Line count should be 4, got", ctx.LineCount)
	}

}
