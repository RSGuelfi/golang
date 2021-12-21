package model

//Imovel representa informações de um imóvel
type Imovel struct {
	X     int    `json:"cordenada_x"`
	Y     int    `json:"cordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

//SetValor define o valor do imovel
func (i *Imovel) SetValor(valor int) {
	i.valor = valor
}

//GetValor retorna o valor do imovel
func (i *Imovel) GetValor() int {
	return i.valor
}
