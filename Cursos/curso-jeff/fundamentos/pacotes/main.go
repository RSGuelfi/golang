package main

import (
	"curso-go-jeff/fundamentos/pacotes/operadora"
	"curso-go-jeff/fundamentos/pacotes/prefixo"
	"fmt"
)

//NomeDOUsuario é o nome do usuário do sistema
var NomeDoUsuario = "Jeff"

func main() {
	fmt.Printf("Nome do usuario: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da operadora: %s\r\n", operadora.NomeDaOperadora)
}
