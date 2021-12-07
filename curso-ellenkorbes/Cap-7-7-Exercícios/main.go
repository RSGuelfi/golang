package main

import "fmt"

func main() {
	clima := 1
	if clima == 1 {
		fmt.Println("chuva")
	} else if clima == 2 {
		fmt.Println("neve")
	} else {
		fmt.Println("sol")
	}
}
