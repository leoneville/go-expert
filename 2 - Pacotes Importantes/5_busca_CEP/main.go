package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, url := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/" + url + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stdout, "Erro ao ler resposta: %v\n", err)
		}
		var data ViaCEP

		if err = json.Unmarshal(res, &data); err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer unmarshal: %v\n", err)
		}

		file, err := os.Create("cidade.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
		}
		defer file.Close()

		_, err = fmt.Fprintf(
			file, "CEP: %s, Localidade: %s, UF: %s", data.Cep, data.Localidade, data.UF,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo: %v\n", err)
		}
		fmt.Println("Arquivo criado com sucesso.")
		fmt.Printf(
			"CEP: %s, Localidade: %s, UF: %s", data.Cep, data.Localidade, data.UF,
		)
	}
}
