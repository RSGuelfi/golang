package main

import "fmt"

var y = 10

func main() {
	z := 21
	qualquercoisa(z)
}

func qualquercoisa(x int) {
	fmt.Println(y)
	fmt.Println(x)
}
