package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	slice[0], slice[1] = 10, 20

	slice = append(slice, 60)
	fmt.Println(slice, len(slice), cap(slice))
}
