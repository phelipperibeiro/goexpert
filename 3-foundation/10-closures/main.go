package main

import (
	"fmt"
)

/*
 * Example 1
 */
func contador() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

/*
 * Example 2
 */
func filtro(p []int, condicao func(int) bool) []int {
	var filtrados []int
	for _, v := range p {
		if condicao(v) {
			filtrados = append(filtrados, v)
		}
	}
	return filtrados
}

func main() {

	/*
	 * Example 1
	 */
	proximoNumero := contador()
	fmt.Println(proximoNumero()) // Output: 1
	fmt.Println(proximoNumero()) // Output: 2
	fmt.Println(proximoNumero()) // Output: 3

	/*
	 * Example 2
	 */
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pares := filtro(numeros, func(x int) bool { return x%2 == 0 })
	fmt.Println(pares) // Output: [2 4 6 8 10]
}

// go run 3-foundation/9.1-closures/*.go
