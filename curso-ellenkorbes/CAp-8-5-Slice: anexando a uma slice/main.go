package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	slice2 := []int{-2, -1, 0}

	slice = append(slice2, slice...)
	fmt.Println(slice)
}
