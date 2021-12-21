package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (abc pessoa) demonstrar() {
	fmt.Println("O nome completo dessa pessoa é:", abc.nome, abc.sobrenome, ",\nEssa pessoa tem", abc.idade, "anos.")
}

func main() {
	humano := pessoa{
		nome:      "Doze",
		sobrenome: "Treze",
		idade:     14,
	}
	humano.demonstrar()
}
