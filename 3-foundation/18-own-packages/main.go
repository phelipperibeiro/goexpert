package main

import (
	"curso-go/mathematics"
	"curso-go/vehicles"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	/*
	 * own package
	 */
	math := mathematics.Soma(10, 20)
	fmt.Println("Resultado: ", math)
	fmt.Println(mathematics.A)

	/*
	 * own package
	 */
	carro := vehicles.Carro{Marca: "Fiat"}
	fmt.Println(carro.Andar())

	/*
	 * external package
	 */
	fmt.Println(uuid.New())
}

// go run 3-foundation/18-own-packages/*.go

// ##################################
// To create a package:
// ##################################

// 1. create a folder, ex: mathematics
// 2. create a file, ex: mathematics.go in the mathematics folder
// 3. create function, ex: func Soma(a int, b int) int { return a + b }
// 4. go init mod <nome-do-modulo>, ex: go mod init curso-go
// 5. import the package, ex: import "curso-go/mathematics"
// 6. call the function, ex: s := mathematics.Soma(10, 20)
// 7. run the main.go file, ex: go run main.go
// 8. go build
