package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var number uint64 = 0

func worker(workerId int, data <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for x := range data {
		atomic.AddUint64(&number, 1)
		fmt.Printf("Worker %d received %d\n", workerId, x+1)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	wg := sync.WaitGroup{}
	qtdWorkers := 1_000_000

	for i := range qtdWorkers {
		wg.Add(1)
		go worker(i, data, &wg)
	}

	for i := range 10_000_000 {
		data <- i
	}
	close(data)

	wg.Wait()
	fmt.Println(number)
}
