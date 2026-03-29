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
	var meuArray [3]int = [3]int{4, 5, 6}
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Println(meuArray)
	fmt.Println(meuArray[len(meuArray)-1])

	for index, value := range meuArray {
		fmt.Printf("O valor do indice %d é %d\n", index, value)
	}
}
