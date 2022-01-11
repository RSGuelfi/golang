package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[12345678] = "Maria"
	aprovados[21353455] = "Pedro"
	aprovados[14475676] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678])
	delete(aprovados, 12345678)      // Maria foi apagada
	fmt.Println(aprovados[12345678]) // Não mostrará nada pois o elemento não existe
}
