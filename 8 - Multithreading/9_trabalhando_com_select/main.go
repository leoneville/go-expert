package main

import (
	"fmt"
	"math/rand/v2"
	"sync/atomic"
	"time"
)

type Message struct {
	ID  uint64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)

	var i uint64

	go func() {
		for {
			atomic.AddUint64(&i, 1)
			msg := Message{i, "Hello from RabbitMQ"}
			time.Sleep(time.Second * time.Duration(rand.Int32N(7)))
			c1 <- msg
		}
	}()

	go func() {
		for {
			atomic.AddUint64(&i, 1)
			msg := Message{i, "Hello from Kafka"}
			time.Sleep(time.Second * time.Duration(rand.Int32N(7)))
			c2 <- msg
		}
	}()

	for {
		select {
		case rabbitMQ := <-c1:
			fmt.Printf("received ID %d from %s\n", rabbitMQ.ID, rabbitMQ.Msg)

		case kafka := <-c2:
			fmt.Printf("received ID %d from %s\n", kafka.ID, kafka.Msg)
		}
	}
}
