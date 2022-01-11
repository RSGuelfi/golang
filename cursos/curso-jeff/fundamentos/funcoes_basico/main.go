package main

import (
	"curso-go-jeff/fundamentos/funcoes_basico/matematica"
	"fmt"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Printf("Resultado da multiplicação = %d\r\n", resultado)
	resultado = matematica.Soma(3, 3)
	fmt.Printf("Resultado da soma = %d\r\n", resultado)
	resultado = matematica.Calculo(matematica.Divisor, 12, 3)
	fmt.Printf("Resultado da divisão = %d\r\n", resultado)
	resultado, resto := matematica.DivisorComResto(12, 5)
	fmt.Printf("Resultado da divisão = %d Resto da divisão %d\r\n", resultado, resto)
}
