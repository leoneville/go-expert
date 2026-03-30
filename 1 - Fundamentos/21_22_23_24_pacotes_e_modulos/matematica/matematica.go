package matematica

import "fmt"

func Soma[T int | float64](a, b T) T {
	return a + b
}

var A int = 10

type Carro struct {
	Marca  string
	Modelo string
}

func (c *Carro) Andar() string {
	return fmt.Sprintf("O %s %s está andando...", c.Marca, c.Modelo)
}
