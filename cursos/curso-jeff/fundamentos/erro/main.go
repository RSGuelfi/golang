package main

import (
	"curso-go-jeff/fundamentos/erro/model"
	"encoding/json"
	"fmt"
)

func main() {

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 17
	casa.Y = 29
	if err := casa.SetValor(1100000000); err != nil {
		fmt.Println("[main] Houve um erro na atribuição de valor a 	casa ", err.Error)
		if err == model.ErrValorDeImovelMuitoAlto {
			fmt.Println("Corretor, por favor verifique a sua avaliação")
		}
		return
	}

	fmt.Printf("Casa é: %+v\r\n", casa)

	objJSON, err := json.Marshal(casa)
	if err != nil {
		fmt.Println("[main] Houve um erro na geração do objeto JSON ", err.Error())
		return
	}
	fmt.Println("A casa em JSON: ", string(objJSON))
}
