package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

const quantidadeDeGoroutines = 600

var contador int

func main() {
	criarGoroutine(quantidadeDeGoroutines)
	wg.Wait()
	fmt.Println("Total de Goroutines", quantidadeDeGoroutines, "\nTotal do contador:", contador)
}

func criarGoroutine(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()

	}
}
