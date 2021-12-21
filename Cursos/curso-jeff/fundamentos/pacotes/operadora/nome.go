package operadora

import (
	"curso-go-jeff/fundamentos/pacotes/prefixo"
	"strconv"
)

//NomeDaOperadora representaa o nome da operadora que se estar a usar
var NomeDaOperadora = "XPTO Telecom"

//PrefixoDaCapitalOperadora prefixo mais o nome da operadora
var PrefixoDaCapitalOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeDaOperadora
