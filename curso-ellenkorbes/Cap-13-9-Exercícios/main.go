package main

import "fmt"

func main() {

	outrafuncao(argumento)
}

func argumento() {
	fmt.Println("outra funcao")
}

func outrafuncao(x func()) {
	fmt.Println("Uma:")
	x()

}
