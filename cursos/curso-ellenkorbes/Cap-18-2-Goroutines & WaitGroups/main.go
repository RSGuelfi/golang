package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	wg.Add(2)
	go func1()
	go func2()

	wg.Wait()
}

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Println("func1:", i)
		time.Sleep(20)
	}
	wg.Done()
}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("func2:", i)
	}
	wg.Done()
}
