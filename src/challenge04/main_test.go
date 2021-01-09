package main

import (
	"testing"
)

func TestSoma(t *testing.T) {
	valor := soma(3, -4)
	if valor != -1 {
		t.Errorf("soma(3, -4) should be -1 but got %d", valor)
	}
}
