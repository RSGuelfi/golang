package main

import (
	"fmt"
	"testing"
)

type test struct {
	data   []int
	answer int
}

func ExampleSoma() {
	fmt.Println(Soma(4, 2, 1))
	// Output:
	// 7
	// 8
	// 9
}

func TestSoma(t *testing.T) {
	x := Soma(3, 2, 1)
	resultado := 6
	if x != resultado {
		t.Error("Expected:", resultado, ", Got:", x)
	}
}

func TestSoma2(t *testing.T) {
	x := Soma(3, 2, 1)
	resultado := 7
	if x != resultado {
		t.Error("Expected:", resultado, ", Got:", x)
	}
}

func TestMultiplicacao(t *testing.T) {
	x := Multiplicacao(10, 10)
	resultado := 100
	if x != resultado {
		t.Error("Expected:", resultado, ", Got:", x)
	}
}

func TestSomaEmTabela(t *testing.T) {
	tests := []test{
		{data: []int{1, 2, 3}, answer: 6},
		{[]int{11, 10, 12}, 34},
		{[]int{-5, 0, 5, 10}, 11},
	}
	for _, v := range tests {
		z := Soma(v.data...)
		if z != v.answer {
			t.Error("Expected:", v.answer, ", Got:", z)
		}
	}
}

func BenchmarkSoma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Soma(1, 1)
	}
}

func BenchmarkMultiplicacao(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiplicacao(2, 1)
	}
}

func BenchmarkSomaEmTabela(b *testing.B) {
	tests := []test{
		{data: []int{1, 2, 3}, answer: 6},
		{[]int{11, 10, 12}, 33},
		{[]int{-5, 0, 5, 10}, 10},
	}
	for i := 0; i < b.N; i++ {
		for _, v := range tests {
			Soma(v.data...)
		}
	}
}
