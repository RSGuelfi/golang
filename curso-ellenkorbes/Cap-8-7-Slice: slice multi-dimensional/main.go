package main

import "fmt"

func main() {
	slices := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(slices[1][1])
}
