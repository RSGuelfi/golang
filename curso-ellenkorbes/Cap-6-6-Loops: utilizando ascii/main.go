package main

import "fmt"

func main() {
	for x := 0; x <= 9999; x++ {
		fmt.Printf("%d - %v\n", x, string(rune(x)))
	}
}
