package main

import "fmt"

func main() {
	canal := make(chan int, 2)
	canal <- 43
	canal <- 42
	fmt.Println(<-canal)
	fmt.Println(<-canal)
}
