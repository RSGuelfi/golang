package main

import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	traçãoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {
	carrãodotio := sedan{veiculo{4, "azul"}, true}
	fubicadovo := caminhonete{
		veiculo: veiculo{
			portas: 6,
			cor:    "ferrugem",
		},
		traçãoNasQuatro: false,
	}

	fmt.Println(carrãodotio)
	fmt.Println(fubicadovo)
}
