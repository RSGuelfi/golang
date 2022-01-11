package main

import "fmt"

func main() {
	hoje := 1

	switch {
	case hoje == 1:
		fmt.Println("1")
	case hoje == 2:
		fmt.Println("2")
	case hoje == 3:
		fmt.Println("3")
	}
}
