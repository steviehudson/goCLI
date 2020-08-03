package main

import (
	"testing"
)

//tests should be suffixed with _test as compiler will ignore them
//json data can be stored in a directory testdata and compiler will ignore

//"go test goCLI" to run tests
//"go test myCLi/..."" to include child packages

//"go testflag help" testflags
//"go test -list Basics goCLI" list tests
//"go test -run Basic goCLI/..." run tests tha match regex passed in
//"go test -timeout .."  use timeout for tests
//"go test -v goCLI" verbose - includes feedback
//"go test -count 10 goCLI/..." run multiple times. used for endurance test or to test remote servers


//"go test -cover goCLI" *rewrites source code so can miss write line numbers
//"go test -coverpkg goCLI" code cover includes standard libraries
//"go test -coverprofile cover.out goCLI" output file
//"go tool cover -html=cover.out" coverage report in html
//"go test -covermode count -coverprofile cover.out goCLI" coverage scale

func Test_BasicChecks(t *testing.T) {
	//run tests in parallel (don't rely on each other)
	t.Parallel()
	t.Run("Go can add", func(t *testing.T) {
		if 1+1 != 2 {
			t.Fail()
		}
	})
	t.Run("Go can concatenate strings", func(t *testing.T) {   
		t.Parallel()
		if "Hello, "+"Go" != "Hello, Go" {
			t.Fail()
		}
	})
}
