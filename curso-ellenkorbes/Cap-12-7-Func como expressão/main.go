package main

import "fmt"

func main() {

	x := 387

	func(x int) {
		fmt.Println(x, "* 8736 =", x*8736)
	}(x)
}
