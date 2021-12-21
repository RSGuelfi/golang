// Package documentacao ilustra como escrever documentação
package documentacao

import "fmt"

// X é um número inútil
const X = 10

// Doc não faz nada
func Doc() {
	fmt.Println("Essa função não faz nada")
}

// Doc2 é a segunda função
func Doc2() {
	fmt.Println("Essa função não faz nada")
}

// doc tem letra minúscula
func doc() {
	fmt.Println("Essa função não faz nada. X é:", X)
}

// Comentário de uma linha.

/*
	Pode ser um comentário de mais de uma linha.
*/
