package main

import (
	"fmt"
)

func main() {
	total := func() int {
		return sum(1, 2, 3, 5, 10, 14) * 2
	}()

	fmt.Println(total)
}

func sum(numeros ...int) int { // Função variadica
	var total int
	for _, numero := range numeros {
		total += numero
	}
	return total
}
