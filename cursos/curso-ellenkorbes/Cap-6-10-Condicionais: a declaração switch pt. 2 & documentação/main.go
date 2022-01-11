package main

import "fmt"

var x interface{}

func main() {
	x = true

	switch x.(type) {
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("bool")
	case string:
		fmt.Println("string")
	case float64:
		fmt.Println("float")
	}
}
