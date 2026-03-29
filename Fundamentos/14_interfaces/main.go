package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
	Ativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado.\n", c.Nome)
}

func (c Cliente) Ativar() {
	c.Ativo = true
	fmt.Printf("O cliente %s foi ativado.\n", c.Nome)
}

func Desativar(pessoa Pessoa) {
	pessoa.Desativar()
}

func Ativar(pessoa Pessoa) {
	pessoa.Ativar()
}

func main() {
	cliente := Cliente{
		Nome:  "Neville",
		Idade: 28,
		Ativo: true,
	}

	cliente.Logradouro = "Aviadores del Chaco"
	cliente.Endereco.Cidade = "Asuncion"
	Desativar(cliente)
	Ativar(cliente)
}
