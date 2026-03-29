package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero uint8   `json:"-"`
	Saldo  float64 `json:"saldo" validate:"gt=0"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string(res))

	if err = json.NewEncoder(os.Stdout).Encode(conta); err != nil {
		panic(err)
	}

	// jsonPuro := []byte(`{"Numero":2, "Saldo":200}`)
	// var contaX Conta

	// if err = json.Unmarshal(jsonPuro, &contaX); err != nil {
	// 	panic(err)
	// }
	// println(contaX.Saldo)

	jsonErrado := []byte(`{"numero":3, "saldo":500}`)
	var contaX2 Conta

	if err = json.Unmarshal(jsonErrado, &contaX2); err != nil {
		panic(err)
	}
	println(contaX2.Saldo)
}
