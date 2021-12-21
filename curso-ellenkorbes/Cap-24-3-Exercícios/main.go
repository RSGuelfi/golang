package main

import "fmt"

type erroEspecial struct {
	qualquercoisa string
}

func (e erroEspecial) Error() string {
	return "!!!"
}

func erroComoParametro(err error) {
	fmt.Println("Erro!:", err.(erroEspecial).qualquercoisa, "\ne no método Erro:", err)
}

func main() {
	arg := erroEspecial{"Erro"}
	erroComoParametro(arg)
}
