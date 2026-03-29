package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	cliente := Cliente{
		Nome:  "Neville",
		Idade: 28,
		Ativo: true,
	}

	cliente.Ativo = false

	fmt.Printf("O cliente %v tem %d e está com status %t", cliente.Nome, cliente.Idade, cliente.Ativo)
}
