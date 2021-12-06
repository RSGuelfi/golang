package main

import (
	"fmt"
	"math"
)

type Motorista struct {
	Latitude       float64
	Longitude      float64
	Name           string
	Available      bool
	CurrentRunning bool
}

type Passageiro struct {
	Latitude  float64
	Longitude float64
	Name      string
}

var valorDir float64

func findMotorista(passageiro []Passageiro, motoristas []Motorista) {
	motoristasAvailable := []Motorista{}
	min := motoristas[0]
	for i, motorista := range motoristas {
		if motoristas[i].Available && !motoristas[i].CurrentRunning {
			motoristasAvailable = append(motoristasAvailable, motoristas[i])
			min = motorista
			tes := min.Distance(passageiro[0])

			defer fmt.Printf("\nMotorista: %v Distance: %.3f Km", i, tes)
		}
	}

	fmt.Println("Motristas Available:", motoristasAvailable)

	// filtrar todos motoristas que estao work_now == false

	//

}

const radius = 6371 // Earth radius in kilometers

func degrees2radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func (origin Motorista) Distance(destination Passageiro) float64 {
	degreesLat := degrees2radians(destination.Latitude - origin.Latitude)
	degreesLong := degrees2radians(destination.Longitude - origin.Longitude)
	a := (math.Sin(degreesLat/2)*math.Sin(degreesLat/2) +
		math.Cos(degrees2radians(origin.Latitude))*
			math.Cos(degrees2radians(destination.Latitude))*math.Sin(degreesLong/2)*
			math.Sin(degreesLong/2))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	valorDir = radius * c

	return valorDir
}

func main() {

	motoristas := []Motorista{
		{-23.51982, -46.79871, "Rodrigo", false, false},
		{-23.51843, -46.79156, "Alberto", true, false},
		{-23.51365, -46.80507, "Lobato", false, true},
		{-23.51432, -46.78825, "Fabio", true, false},
		{-23.52425, -46.76009, "Leonardo", false, true},
		{-23.51437, -46.76543, "Ricardo", true, false},
		{-23.45432, -45.13256, "Alice", false, false},
		{-23.34567, -46.94542, "Lucas", false, true},
		{-23.67891, -47.41223, "Solange", false, false},
		{-23.37465, -46.83456, "Renato", true, true},
		{-23.51437, -46.76510, "Gabriel", false, false},
	}

	var passageiro = []Passageiro{{-23.51411, -46.76011, "Joao"}}

	findMotorista(passageiro, motoristas)
}

// func DifLatLong(value [10]Motorista) {

// 	minLong := math.Min(float64(motoristas[0].Latitude), float64(motoristas[1].Latitude))
// 	minLati := math.Min(float64(motoristas[0].Longitude), float64(motoristas[1].Longitude))
// 	fmt.Println("minLong: ", minLong, "\nminLati: ", minLati)
// }

// func tstminval(arr []float64) float64 {
// 	min := arr[0] // assume first value is the smallest

// 	for _, value := range arr {
// 		if value < min {
// 			min = value // found another smaller value, replace previous value in min
// 		}
// 	}
// 	return min
// }

// func findMinElement(arr []float64) float64 {
// 	min_num := arr[0]

// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] < min_num {
// 			min_num = arr[i]
// 		}
// 	}
// 	return min_num
// }
