package main

import "fmt"

func main() {
	x := 20

	// valor(x)
	ponteiro(&x)
	fmt.Println(x)
}

func valor(x int) {
	x++
	fmt.Println(x)
}

func ponteiro(x *int) {
	*x++
	fmt.Println(x)
}
