package main

import (
	"fmt"
	"golib"
	"sync"
)

func main() {
	golib.SayHello()

	//go routines
	//simple example of race conditions, more complicated when go routines shared across applicaitons
	//identify race conditions using "go run -race  main.go"
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		//redefine i within wait group
		i := i
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}
