package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 10 {
		fmt.Printf("%d: Task %s is running\n", i+1, name)
		time.Sleep(1 * time.Second)
	}
}

// Thread 1 (Principal)
func main() {
	waitGroup := sync.WaitGroup{}
	// Thread 2
	waitGroup.Add(1)
	go task("A", &waitGroup)

	// Thread 3
	waitGroup.Add(1)
	go task("B", &waitGroup)

	// Thread 4
	waitGroup.Go(func() {

		for i := range 5 {
			fmt.Printf("%d: Task %s is running\n", i+1, "anonymous")
			time.Sleep(1 * time.Second)
		}
	})

	waitGroup.Wait()
}
