package main

import (
	"testing"
)

func TestAdder(t *testing.T) {
	if adder(2, 5) != 7 {
		t.Fail()
	}
}

//works in real time, certain seconds test attempts to execute
//"go test -bench Adder goCLI" or "go test -bench . goCLI" or "go test -benchtime 5s -bench . goCLI"
//"go test -bench . -benchmem goCLI" with memory allocation

func BenchmarkAdder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		adder(5, 7)
	}
}

//test profiling

//"go test -memprofile mem.out goCLI" -create test profile output
//"go tool pprof -web goCLI.test mem.out" - look for mem leeks
//see https://app.pluralsight.com/course-player?clipId=92c02f66-6f3e-413e-b1ca-a7b362374fbf for further analysis
func TestMessageWriter(t *testing.T) {
	if messageWriter("Hello", "World") != "Hello, World" {
		t.Fail()
	}
}
