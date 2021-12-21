package model

import "errors"

//Imovel representa informações de um imóvel
type Imovel struct {
	X     int    `json:"cordenada_x"`
	Y     int    `json:"cordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

//ErrValorDeImovelInvalido é um erro para imoveis abaixo do preço normal
var ErrValorDeImovelInvalido = errors.New("O valor informado não é valido para um imovel")

//ErrValorDeImovelMuitoAlto é um erro caso o valor for muito alto
var ErrValorDeImovelMuitoAlto = errors.New("O valor informado é muito alto.")

//GetValor retorna o valor do imovel
func (i *Imovel) GetValor() int {
	return i.valor
}

//SetValor define o valor do imovel
func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorDeImovelInvalido
		return
	} else if valor > 1000000 {
		err = ErrValorDeImovelMuitoAlto
		return
	}
	i.valor = valor
	return
}
