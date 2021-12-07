package main

import (
	"fmt"
)

func main() {
	a := map[int]string{
		1: "abc",
		2: "def",
		3: "ghi",
	}

	fmt.Println(a)
	total := 0

	for key, _ := range a {
		total += key
	}

	fmt.Println(total)

	delete(a, 2)
	fmt.Println(a)
}
