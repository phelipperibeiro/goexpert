package main

import (
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	res, err := io.ReadAll(req.Body) // reader
	if err != nil {
		panic(err)
	}
	println(string(res))
	req.Body.Close()
}

// go run 4-important-packages/2-http-requests/main.go
