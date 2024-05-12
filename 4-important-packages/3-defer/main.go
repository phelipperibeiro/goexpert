package main

import "fmt"

func main() {
	defer fmt.Println("Primeira Linha") // executes after the main function ends
	fmt.Println("Segunda Linha")
	fmt.Println("Terceira Linha")
}

// go run 4-important-packages/3-defer/main.go
