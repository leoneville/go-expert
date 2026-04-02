package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/leoneville/math"
)

func main() {
	m := math.NewMath(1, 2)
	print(m.Add())
	fmt.Println(uuid.New().String())
}
