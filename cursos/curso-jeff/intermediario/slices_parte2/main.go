package main

import "fmt"

func main() {
	capitais := []string{"Lisboa"}
	//fmt.Println(capitais, len(capitais), cap(capitais))
	capitais = append(capitais, "Brasília")
	//fmt.Println(capitais, len(capitais), cap(capitais))
	cidades := make([]string, 5)
	cidades[0] = "Nova York"
	cidades[1] = "Londres"
	cidades[2] = "Madeira"
	cidades[3] = "Tokyo"
	cidades[4] = "Singapura"
	fmt.Println(cidades, len(cidades), cap(cidades))
	capitais[1] = "Brasilia"
	fmt.Println(capitais, len(capitais), cap(capitais))
	for indice, cidade := range cidades {
		fmt.Printf("Cidade[%d] = %s\n\r", indice, cidade)
	}
	//Primeiro item começa com indice 0
	//Segundo item começa com indice 1
	capitaisAsia := cidades[3:5]
	fmt.Println(capitaisAsia)
	temp1 := cidades[:2]
	fmt.Println(temp1)
	temp2 := cidades[len(cidades)-2:]
	fmt.Println(temp2)
	indiceDoItemARetirar := 2
	temp := cidades[:indiceDoItemARetirar]
	temp = append(temp, cidades[indiceDoItemARetirar+1:]...)
	copy(cidades, temp)
	fmt.Println(cidades)
}
