package matematica

/*
Calculo executa qualquer tipo de calculo basta enviar a funcao desejada
*/
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

//Multiplicador multiplica x vezes y
func Multiplicador(x int, y int) int {
	return x * y
}

//Divisor efetua a divisão do numeroA com o numeroB
func Divisor(numeroA int, numeroB int) (resultado int) {
	resultado = numeroA / numeroB
	return
}

//DivisorComResto efetua a divisão di numeroA pelo numeroB
//retorna o resultado e também o resto da divisão
func DivisorComResto(numeroA int, numeroB int) (resultado int, resto int) {
	resultado = numeroA / numeroB
	resto = numeroA / numeroB
	return
}
