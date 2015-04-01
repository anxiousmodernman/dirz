package dirz

import (
	"testing"
)

func TestParseSingleFile(t *testing.T) {
	var filename = "./testfiles/touch-these-files.dirz"
	ParseFile(filename)
}
