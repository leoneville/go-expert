package main

import "fmt"

// Thread 1 (Principal)
func main() {
	canal := make(chan string) // Canal Vazio

	// Thread 2
	go func() {
		canal <- "Olá Mundo!" // Canal Cheio
	}()

	// Thread 1
	msg := <-canal // Canal Esvazia
	fmt.Println(msg)
}
