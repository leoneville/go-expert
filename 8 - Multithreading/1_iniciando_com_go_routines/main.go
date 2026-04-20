package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := range 10 {
		fmt.Printf("%d: Task %s is running\n", i+1, name)
		time.Sleep(1 * time.Second)
	}
}

// Thread 1 (Principal)
func main() {
	// Thread 2
	go task("A")

	// Thread 3
	go task("B")

	// Thread 4
	go func() {
		for i := range 10 {
			fmt.Printf("%d: Task %s is running\n", i+1, "anonymous")
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(15 * time.Second)
}
