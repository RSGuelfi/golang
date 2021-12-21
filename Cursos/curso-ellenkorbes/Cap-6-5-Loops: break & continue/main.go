package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}
