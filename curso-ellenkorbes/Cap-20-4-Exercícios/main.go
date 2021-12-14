package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var contador int

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
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()

	}
}
