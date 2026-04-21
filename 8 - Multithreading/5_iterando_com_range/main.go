package main

import (
	"fmt"
)

// Thread 1 (Principal)
func main() {
	canal := make(chan int)

	// Thread 2
	go publish(canal)

	// Thread 3
	reader(canal)
}

func publish(ch chan<- int) {
	defer close(ch)

	for i := range 10 {
		ch <- i
	}
}

func reader(ch <-chan int) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
	}
}
