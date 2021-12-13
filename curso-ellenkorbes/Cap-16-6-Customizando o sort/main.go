package main

import (
	"fmt"
	"sort"
)

type Carro struct {
	Nome     string
	Potencia int
	Consumo  int
}

type ordenarPorPotencia []Carro

func (x ordenarPorPotencia) Len() int           { return len(x) }
func (x ordenarPorPotencia) Less(i, j int) bool { return x[i].Potencia < x[j].Potencia }
func (a ordenarPorPotencia) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ordenarPorConsumo []Carro

func (x ordenarPorConsumo) Len() int           { return len(x) }
func (x ordenarPorConsumo) Less(i, j int) bool { return x[i].Consumo > x[j].Consumo }
func (a ordenarPorConsumo) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ordenarPorLucroProDonoDoPosto []Carro

func (x ordenarPorLucroProDonoDoPosto) Len() int           { return len(x) }
func (x ordenarPorLucroProDonoDoPosto) Less(i, j int) bool { return x[i].Consumo < x[j].Consumo }
func (a ordenarPorLucroProDonoDoPosto) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {

	carros := []Carro{{"Chevete", 50, 8},
		{"Porsche", 300, 5},
		{"Fusca", 20, 30},
	}

	fmt.Println("Inicial:\n", carros)

	sort.Sort(ordenarPorPotencia(carros))

	fmt.Println("Potência:\n", carros)

	sort.Sort(ordenarPorConsumo(carros))

	fmt.Println("Consumo:\n", carros)

	sort.Sort(ordenarPorLucroProDonoDoPosto(carros))

	fmt.Println("Lucro:\n", carros)

}
