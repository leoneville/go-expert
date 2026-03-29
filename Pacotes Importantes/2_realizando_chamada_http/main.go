package main

import (
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://viacep.com.br/ws/60861212/json/")
	if err != nil {
		panic(err)
	}

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))
	if err = req.Body.Close(); err != nil {
		panic(err)
	}
}
