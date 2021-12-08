package main

import "fmt"

func main() {
	x := struct {
		nome  string
		idade int
	}{
		nome:  "mario",
		idade: 50,
	}
	fmt.Println(x)

}
