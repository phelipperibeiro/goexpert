package main

import "fmt"

/*
 * ###################################################################
 * function that receives a pointer to an integer and increments it
 * ###################################################################
 */
func incrementa(valor *int) {
	*valor = *valor + 1 // increment the value pointed by the pointer
}

func main() {
	/*
	 * Example 1
	 */
	var x int = 10
	var ponteiro *int // declaration of a pointer to an integer
	ponteiro = &x     // assigns the memory address of x to the pointer

	fmt.Println("Valor de x:", x)
	fmt.Println("Endereço de x:", &x)
	fmt.Println("Valor apontado pelo ponteiro:", *ponteiro) // dereferencing the pointer to access the value of x

	*ponteiro = 20 // changing the value of x through the pointer
	fmt.Println("Novo valor de x:", x)

	/*
	 * Example 2
	 */
	numero := 10
	fmt.Println("Valor original:", numero)

	// Passing the memory address of 'numero' to the incrementa function
	incrementa(&numero)
	fmt.Println("Valor após a incrementação:", numero)

}

// go run 3-foundation/12-pointers/*.go
