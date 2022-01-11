package main

import "fmt"

var i int
var f float64
var s string
var b bool

func main() {
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", f, f)
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", b, b)
}
