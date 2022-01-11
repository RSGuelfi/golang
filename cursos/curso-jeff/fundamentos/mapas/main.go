package main

import (
	"curso-go-jeff/fundamentos/struct_avancado/model"
	"fmt"
)

var cache map[string]model.Imovel

func main() {

	cache = make(map[string]model.Imovel)

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(600000)

	cache["Casa Amarela"] = casa

	fmt.Println("Há uma casa amarela no cache?")
	imovel, achou := cache["Casa Amarela"]
	if achou {
		fmt.Printf("Sim, e o que achei foi: %+v\r\n", imovel)
	}

	apto := model.Imovel{}
	apto.Nome = "Apto Azul"
	apto.X = 19
	apto.Y = 26
	apto.SetValor(700000)

	cache[apto.Nome] = apto

	fmt.Println("Quantos itens há no cache? ", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("Chave[%s] = %+v\r\n", chave, imovel)
	}

	imovel, achou = cache["Casa Amarela"]
	if achou {
		delete(cache, imovel.Nome)
	}

	fmt.Println("Quantos itens há no cache ", len(cache))
}
