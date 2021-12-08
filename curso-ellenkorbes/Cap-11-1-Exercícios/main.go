package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {

	pessoa1 := pessoa{
		nome:      "Renata",
		sobrenome: "Pimeta",
		sabores:   []string{"pistache", "baunilha", "morango"},
	}
	fmt.Println(pessoa1.nome, ":")
	for _, v := range pessoa1.sabores {
		fmt.Println("\t", v)
	}
}
