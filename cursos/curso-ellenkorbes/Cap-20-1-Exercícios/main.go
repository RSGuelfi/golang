package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	novagoroutine(100)
	wg.Wait()
}

func novagoroutine(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		x := j
		go func(i int) {
			fmt.Println("Goroutine 1:", i+1)
			wg.Done()
		}(x)
	}
}
