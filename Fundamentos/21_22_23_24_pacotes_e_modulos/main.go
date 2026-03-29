package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/leoneville/curso-go/matematica"
)

func main() {
	s := matematica.Soma(10, 50)
	carro := matematica.Carro{Marca: "Renault", Modelo: "Niagara"}
	fmt.Println(carro.Andar())
	fmt.Printf("Resultado: %v\n", s)
	fmt.Println(matematica.A)

	uuid7, err := uuid.NewV7()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(uuid7)
}
