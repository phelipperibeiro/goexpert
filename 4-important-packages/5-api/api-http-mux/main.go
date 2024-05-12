package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World!"))
	// })
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "My Blog"})
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}

// go run 4-important-packages/5-api/api-http-mux/main.go
