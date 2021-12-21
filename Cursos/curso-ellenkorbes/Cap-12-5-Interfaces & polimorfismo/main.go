package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type dentista struct {
	pessoa
	dentesarrancados int
	salario          string
}

type arquiteto struct {
	pessoa
	tipodeconstrução string
	tamanhodaloucura string
}

func (x dentista) oibomdia() {
	fmt.Println("Meu nome é", x.nome, "e eu já arranquei", x.dentesarrancados, "dentes, e ouve só: Bom dia!")
}

func (x arquiteto) oibomdia() {
	fmt.Println("Meu nome é", x.nome, "e ouve só: Bom dia!")
}

type gente interface {
	oibomdia()
}

func serhumano(g gente) {
	g.oibomdia()
	switch g.(type) {
	case dentista:
		fmt.Println("Eu ganho:", g.(dentista).salario)
	case arquiteto:
		fmt.Println("Eu construo:", g.(arquiteto).tipodeconstrução)
	}
}

func main() {
	mrdente := dentista{
		pessoa: pessoa{
			nome:      "Mister Dente",
			sobrenome: "da Silva",
			idade:     50,
		},
		dentesarrancados: 8973,
		salario:          "9.787",
	}
	mrpredio := arquiteto{
		pessoa: pessoa{
			nome:      "Mister Predio",
			sobrenome: "da Silva",
			idade:     59,
		},
		tipodeconstrução: "Hospícios",
		tamanhodaloucura: "Demais",
	}
	serhumano(mrdente)
	fmt.Println("")
	serhumano(mrpredio)

}
