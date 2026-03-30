package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"
)

type UsuarioRequest struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func main() {
	client := http.Client{Timeout: time.Second * 30}
	usuario := UsuarioRequest{
		Nome:  "Neville",
		Email: "email@email.com",
	}

	jsonBytes, err := json.Marshal(usuario)
	if err != nil {
		panic(err)
	}
	jsonVar := bytes.NewBuffer(jsonBytes)

	resp, err := client.Post("http://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
