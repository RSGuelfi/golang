package main

import "fmt"

func main() {
	x := 1.1
	for x < 10 {
		fmt.Printf("chis é menor que déis %.2f\n", x)
		x += 1.1
		if x >= 10 {
			fmt.Printf("chis é maior que déis %.2f\n", x)

		}
	}
}
