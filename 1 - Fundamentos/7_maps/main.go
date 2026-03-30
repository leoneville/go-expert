package main

import "fmt"

func main() {
	salarios := map[string]int{"Neville": 1500, "Lais": 2000}
	fmt.Println(salarios)
	delete(salarios, "Neville")
	salarios["Neville"] = 5000
	fmt.Println(salarios["Neville"])

	salario := make(map[string]int)
	salario["Teste"] = 123

	salario_1 := map[string]int{}
	salario_1["Teste"] = 123

	for chave, valor := range salarios {
		fmt.Printf("O salário do %s é de %d\n", chave, valor)
	}
}
