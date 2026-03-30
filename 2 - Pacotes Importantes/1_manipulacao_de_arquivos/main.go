package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Escrita de arquivo
	f, err := os.OpenFile("arquivo.txt", os.O_APPEND, os.ModeAppend)
	if err != nil {
		log.Println("Criando arquivo.txt")
		f, err = os.Create("arquivo.txt")
		if err != nil {
			panic(err)
		}
	}

	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo!\n"))
	// tamanho, err := f.WriteString("Hello World from File\n")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)
	f.Close()

	// Atualizar arquivo
	f2, err := os.OpenFile("arquivo.txt", os.O_APPEND, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	f2.Write([]byte("Algum texto adicionado!\n"))

	f2.Close()

	// Leitura de arquivo
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	// leitura de pouco em pouco abrindo o arquivo
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(n)
		fmt.Println(string(buffer[:n]))
	}

	arquivo2.Close()

	if err = os.Remove("arquivo.txt"); err != nil {
		panic(err)
	}
}
