package main

type Conta struct {
	saldo int
	Ativa bool
}

func NewConta() *Conta {
	return &Conta{saldo: 300, Ativa: true}
}

func (c *Conta) simular(saldo int) int {
	c.saldo += saldo
	println(c.saldo)
	return c.saldo
}

func main() {
	conta := NewConta()

	conta.simular(200)
	println(conta.saldo)
}
