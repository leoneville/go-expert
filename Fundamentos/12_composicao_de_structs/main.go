package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	cliente := Cliente{
		Nome:  "Neville",
		Idade: 28,
		Ativo: true,
	}

	cliente.Ativo = false
	cliente.Logradouro = "Aviadores del Chaco"
	cliente.Endereco.Cidade = "Asuncion"

	fmt.Printf("O cliente %v tem %d e está com status %t", cliente.Nome, cliente.Idade, cliente.Ativo)
}
