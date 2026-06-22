package main

import "github.com/leoneville/fcutils/pkg/rabbitmq"

func main() {
	ch := rabbitmq.OpenChannel()
	defer ch.Close()

	rabbitmq.Publish(ch, "Hello World!", "amq.direct")
}
