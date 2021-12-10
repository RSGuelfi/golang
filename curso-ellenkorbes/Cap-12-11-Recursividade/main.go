package main

import "fmt"

func main() {
	fmt.Println(fatorial(10))
	fmt.Println(loops(10))
}

func fatorial(x int) int {
	if x == 1 {
		return x
	}
	return x * fatorial(x-1)
}

func loops(x int) int {
	total := 1
	for x > 1 {
		total *= x
		x--
	}
	return total
}
