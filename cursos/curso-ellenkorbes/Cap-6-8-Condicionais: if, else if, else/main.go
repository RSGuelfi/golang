package main

import "fmt"

func main() {
	if x := 10; x > 100 {
		fmt.Println("Maior que cem")
	} else if x < 10 {
		fmt.Println("Menor que cem")
	} else {
		fmt.Println("Diferente")
	}
}
