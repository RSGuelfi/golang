package main

import "fmt"

func main() {
	c := make(chan int)
	go meuloop(10, c)

	prints(c)
}

func meuloop(t int, s chan<- int) {
	for i := 0; i < t; i++ {
		s <- i
	}
	close(s)
}

func prints(r <-chan int) {
	for v := range r {
		fmt.Println("Recebido:", v)
	}
}
