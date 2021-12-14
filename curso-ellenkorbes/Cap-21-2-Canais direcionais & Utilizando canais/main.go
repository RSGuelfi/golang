package main

import "fmt"

func main() {
	canal := make(chan int)

	go send(canal)

	receive(canal)

}

func send(s chan<- int) {
	s <- 42
}

func receive(c <-chan int) {
	fmt.Println("Valor Recebido:", <-c)
}
