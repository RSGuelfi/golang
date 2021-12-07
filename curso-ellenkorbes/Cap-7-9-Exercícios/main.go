package main

import "fmt"

func main() {
	semana := "segunda"

	switch semana {
	case "domingo":
		fmt.Println("domingo")
	case "segunda":
		fmt.Println("segunda")
	case "terça":
		fmt.Println("terça")
	case "quarta":
		fmt.Println("quinta")
	case "sexta":
		fmt.Println("sexta")
	case "sabado":
		fmt.Println("sabado")
	}
}
