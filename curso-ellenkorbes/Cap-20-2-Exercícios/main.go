package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func (p *Pessoa) falar() {
	fmt.Println(p.Nome, "Diz oi!")
}

type Humanos interface {
	falar()
}

func dizerAlgumaCOisa(h Humanos) {
	h.falar()
}

func main() {

	p1 := Pessoa{"Genghis Khan", 1000}

	p1.falar()

	(&p1).falar()

	dizerAlgumaCOisa(&p1)
}
