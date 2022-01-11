package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var mu sync.Mutex
var contador int32

const quantidadeDeGoroutines = 600

func main() {
	criarGoroutine(quantidadeDeGoroutines)
	wg.Wait()
	fmt.Println("Total de Goroutines", quantidadeDeGoroutines, "\nTotal do contador:", contador)
}

func criarGoroutine(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			atomic.AddInt32(&contador, 1)
			v := atomic.LoadInt32(&contador)
			runtime.Gosched()
			fmt.Println(v)
			wg.Done()
		}()

	}
}
