package main

import "fmt"

func main() {
	x := 10

	y := &x

	fmt.Println(x, y)

	*y++

	fmt.Println(*y)
	fmt.Printf("%T, %T\n", x, y)
	fmt.Println(x, y)
}
