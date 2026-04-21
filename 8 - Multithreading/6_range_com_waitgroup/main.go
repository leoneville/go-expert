package main

import (
	"fmt"
	"sync"
)

// Thread 1 (Principal)
func main() {
	canal := make(chan int)
	wg := sync.WaitGroup{}

	go publish(canal)

	wg.Add(1)
	go reader(canal, &wg)

	wg.Wait()
}

func publish(ch chan<- int) {
	defer close(ch)

	for i := range 10 {
		ch <- i
	}
}

func reader(ch <-chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
	}

	wg.Done()
}
