package main

import "fmt"

func main() {
	slices := [][]string{
		[]string{"abc", "def", "ghi"},
		[]string{"ghi", "def", "abc"},
	}
	for _, v := range slices {
		fmt.Println(v)
	}
}
