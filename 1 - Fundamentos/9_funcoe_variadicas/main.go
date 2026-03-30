package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2, 3, 5, 10, 14))
}

func sum(numeros ...int) int { // Função variadica
	var total int
	for _, numero := range numeros {
		total += numero
	}
	return total
}
