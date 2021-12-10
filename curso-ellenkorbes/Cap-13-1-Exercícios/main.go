package main

import "fmt"

func main() {
	fmt.Println(retornaint(4))
	fmt.Println(retornaintestring(8, "oito"))
}

func retornaint(x int) int {
	return x
}

func retornaintestring(x int, y string) (int, string) {
	return x, y
}
