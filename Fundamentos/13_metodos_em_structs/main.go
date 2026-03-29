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

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado.", c.Nome)
}

func main() {
	cliente := Cliente{
		Nome:  "Neville",
		Idade: 28,
		Ativo: true,
	}

	cliente.Logradouro = "Aviadores del Chaco"
	cliente.Endereco.Cidade = "Asuncion"
	cliente.Desativar()

}
