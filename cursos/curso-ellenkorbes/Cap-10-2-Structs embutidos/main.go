package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario int
}

func main() {
	pessoa1 := pessoa{
		nome:  "Alfredo",
		idade: 30,
	}
	pessoa2 := profissional{
		pessoa: pessoa{nome: "Maricota",
			idade: 30},
		titulo:  "Pizzaiolo",
		salario: 1300,
	}
	pessoa3 := pessoa{nome: "Mauricio", idade: 35}
	pessoa4 := profissional{pessoa: pessoa{nome: "Vanderlei", idade: 60}, titulo: "Politico", salario: 13000}
	fmt.Println(pessoa1.idade, pessoa1.nome)
	fmt.Println(pessoa2)
	fmt.Println(pessoa3)
	fmt.Println(pessoa4)
}
