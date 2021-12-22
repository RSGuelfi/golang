package main

import (
	"fmt"
	"testing"
)

type test struct {
	data   []int
	answer int
}

func TestSomaEmTabela(t *testing.T) {
	tests := []test{
		{data: []int{1, 2, 3}, answer: 6},
		{[]int{11, 10, 12}, 33},
		{[]int{-5, 0, 5, 10}, 10},
	}
	for _, v := range tests {
		z := Soma(v.data...)
		if z != v.answer {
			t.Error("Expected:", v.answer, "Got:", z)
		}
	}
}

func ExampleSoma() {
	fmt.Println(Soma(4, 2, 1))
	fmt.Println(Soma(4, 2, 2))
	fmt.Println(Soma(4, 2, 3))
	// Output:
	// 7
	// 8
	// 9
}

func TestSoma(t *testing.T) {
	x := Soma(3, 2, 1)
	resultado := 6
	if x != resultado {
		t.Error("Expected:", resultado, "Got:", x)
	}
}

func TestMultiplicacao(t *testing.T) {
	x := Multiplicacao(10, 10)
	resultado := 100
	if x != resultado {
		t.Error("Expected:", resultado, ", Got:", x)
	}
}
