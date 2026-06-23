package main

import (
	"fmt"

	"github.com/leoneville/fcutils-secret/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
