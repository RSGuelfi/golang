package main

import "fmt"

func main() {
	for horas := 0; horas <= 12; horas++ {
		fmt.Println("Mẽs: ", horas)
		for x := 0; x <= 30; x++ {
			fmt.Print("Dia: ", x, ", ")
		}
		fmt.Println("\n")
	}
}
