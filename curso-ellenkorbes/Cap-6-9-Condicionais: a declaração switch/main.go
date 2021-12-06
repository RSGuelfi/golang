package main

import "fmt"

func main() {
	x := 1

	switch true {
	case (x == 4), (x == 8):
		fmt.Println("x é 4 ou 8")
	case (x < 10):
		fmt.Println("menor que 10")
	default:
		fmt.Println(0)
	}
}
