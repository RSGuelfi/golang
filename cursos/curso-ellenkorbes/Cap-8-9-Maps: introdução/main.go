package main

import "fmt"

func main() {
	a := map[int]int{
		1: 2,
		2: 4,
		3: 6,
	}

	fmt.Println(a)
	fmt.Println(a[1])

	a[3] = 3 + 6

	fmt.Println(a)
}
