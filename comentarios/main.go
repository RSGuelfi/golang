package main

import (
	"fmt"
	"sort"
)

func main() {

	vals := []float64{3, 1, 0, 7, 2, 4, 8, 6}
	sort.Float64s(vals)
	fmt.Println(vals)

	sort.Sort(sort.Reverse(sort.Float64Slice(vals)))
	fmt.Println(vals)
}
