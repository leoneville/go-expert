package main

import "fmt"

const a = "Hello world"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Neville"
	e float64 = 15.4
	f ID      = 1
)

func main() {
	fmt.Printf("O tipo de E é %T\n", e)
	fmt.Printf("O tipo de E é %T", f)
}
