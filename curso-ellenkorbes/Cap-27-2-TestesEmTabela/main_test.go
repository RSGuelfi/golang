package main

import "testing"

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
		z := soma(v.data...)
		if z != v.answer {
			t.Error("Expected:", v.answer, "Got:", z)
		}
	}
}

func TestSoma(t *testing.T) {
	x := soma(3, 2, 1)
	resultado := 6
	if x != resultado {
		t.Error("Expected:", resultado, "Got:", x)
	}
}

func TestSoma2(t *testing.T) {
	x := soma(3, 2, 1)
	resultado := 7
	if x != resultado {
		t.Error("Expected:", resultado, "Got:", x)
	}
}
