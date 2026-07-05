package main

import (
	"fmt"
	"os"
)

func main() {
	i := 0
	for {
		if i >= 1000 {
			break
		}
		f, err := os.Create(fmt.Sprintf("./tmp/file%d.txt", i))
		if err != nil {
			panic(err)
		}

		f.WriteString("Hello, World!")
		f.Close()
		i++
	}
}
