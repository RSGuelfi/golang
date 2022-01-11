package main

import "fmt"

func main() {
	s := "ascii é噇᎔င 蜤"
	for _, v := range s {
		fmt.Printf("%b - %T - %#U - %#X\n", v, v, v, v)
	}

	fmt.Println("")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%b - %T - %#U - %#X\n", s[i], s[i], s[i], s[i])
	}
}
