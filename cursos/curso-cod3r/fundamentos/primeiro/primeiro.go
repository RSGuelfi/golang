// Programas executáveis iniciam pelo pacote main
package main

/*
Os códigos em Go são organizados em pacotes
e para usá-los é necessário declarar um ou vários imports
*/
import "fmt"

// A porta de entrada de um programa Go é a funcão main
func main() {
	fmt.Print("Primeiro ")
	fmt.Print("Programa!")
}