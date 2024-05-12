package main

import "fmt"

func soma(a, b *int) int {
	*a = 50
	*b = 50
	return *a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20

	fmt.Printf("%d\n", soma(&minhaVar1, &minhaVar2))
	fmt.Printf("%d\n", minhaVar1)
	fmt.Printf("%d\n", minhaVar2)

}

// go run 3-foundation/12.1-pointers-use-case/*.go
