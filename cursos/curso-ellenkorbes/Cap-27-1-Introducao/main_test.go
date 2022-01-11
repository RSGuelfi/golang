package main

import "testing"

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
