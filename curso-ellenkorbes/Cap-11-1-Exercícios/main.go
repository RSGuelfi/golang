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
		sabores:   []string{"pistache", "baunilha", ""},
	}
	fmt.Println(pessoa1)
}
