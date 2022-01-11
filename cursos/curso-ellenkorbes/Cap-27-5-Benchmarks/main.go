package main

import "fmt"

func main() {
	x := Soma(1, 2, 3)
	y := Multiplicacao(1, 2, 3)
	fmt.Println(x)
	fmt.Println(y)
}

func Soma(i ...int) int {
	total := 0
	for _, v := range i {
		total = total + v
	}
	return total
}

func Multiplicacao(i ...int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}
	return total
}
