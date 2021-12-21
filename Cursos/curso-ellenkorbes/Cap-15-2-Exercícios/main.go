package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func main() {

	zezinho := pessoa{"Zezinho", "Primodoluizinho", 10}
	fmt.Println(zezinho)
	mudeme(&zezinho)
	fmt.Println(zezinho)
}

func mudeme(p *pessoa) {
	(*p).nome = "Luizinho"
	p.sobrenome = "Primodozezinho"
}
