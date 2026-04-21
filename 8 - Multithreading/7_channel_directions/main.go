package main

import "fmt"

func recebe(nome string, hello chan<- string) { // Receive-Only
	hello <- nome
}

func ler(data <-chan string) { // Send-Only
	fmt.Println(<-data)
}

func main() {
	canal := make(chan string)

	go recebe("Neville", canal)
	ler(canal)
}
