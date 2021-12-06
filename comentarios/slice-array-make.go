package main

import "fmt"

func main() {
	array1 := [7]int{1, 2, 3, 45}
	slice1 := []int{1, 2, 3, 45}
	fmt.Println(array1, slice1)

	slice1 = array1[1:5]
	fmt.Println(slice1)

	slice1 = append(slice1, 101)
	fmt.Println(slice1)

	slice1 = make([]int, 14, 18)
	slice1[6] = 7
	fmt.Println(slice1, len(slice1), cap(slice1))

	slice1 = append(slice1, 15, 16, 17, 18, 19, 20)
	fmt.Println(slice1, len(slice1), cap(slice1))

	slice2 := make([]int, 5, 10)
	slice3 := append(slice2, 5, 10, 15)
	fmt.Println(slice2, slice3)

	slice2[3] = 7
	fmt.Println(slice2, slice3)

	array2 := [3]int{11, 22, 33}
	var slice4 []int
	slice4 = append(slice4, 44, 55, 66)
	fmt.Println(array2, slice4)

	slice5 := make([]int, 2)
	copy(slice5, slice4)
	fmt.Println(slice5)
}
