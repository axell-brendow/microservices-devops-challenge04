package main

import (
	"fmt"
)

func soma(a int, b int) int {
	return a + b
}

func main() {
	valor := soma(5, 5)
	fmt.Println(valor)
}
