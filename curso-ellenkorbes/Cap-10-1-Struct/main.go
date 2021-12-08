package main

import "fmt"

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func main() {
	cliente1 := cliente{
		nome:      "João",
		sobrenome: "da Silva",
		fumante:   false,
	}
	cliente2 := cliente{nome: "Joana", sobrenome: "Pereira", fumante: true}

	fmt.Println(cliente1)
	fmt.Println(cliente2)
}
