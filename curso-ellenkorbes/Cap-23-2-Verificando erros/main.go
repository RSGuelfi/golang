package main

import "fmt"

func main() {
	n, err1 := fmt.Println("Hello")
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(n)
}
