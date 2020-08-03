package main

import (
	"fmt"
	"golib"
)

func main() {
	golib.SayHello()
}

func adder(l, r int) int {
	return l + r
}

func subractor(l, r int) int {
	return l - r
}

func messageWriter(greeting, name string) string {
	message := fmt.Sprintf("%v, %v", greeting, name)
	return message
}
