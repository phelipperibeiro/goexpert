package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("Iniciando servidor na porta :8080")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")
	select {
	case <-time.After(5 * time.Second):
		// Imprime no comand line stdout
		log.Println("Request processada com sucesso")
		// Imprime no browser
		response.Write([]byte("Request processada com sucesso"))
	case <-ctx.Done():
		// Imprime no comand line stdout
		log.Println("Request cancelada pelo cliente")
	}
}

// go run 5-context/2-client-Server/server/main.go
