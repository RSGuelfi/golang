package main

import "fmt"

func main() {
	x := retornafuncao()

	x()
}

func retornafuncao() func() {
	return func() {
		fmt.Println("outra funcao")
	}
}
