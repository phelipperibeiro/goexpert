package main

import "fmt"

func main() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Println(meuArray[0])               // first value
	fmt.Println(meuArray[len(meuArray)-1]) // last value
	fmt.Println(len(meuArray) - 1)         // last index
	fmt.Println(len(meuArray))             // length

	for i, v := range meuArray {
		fmt.Println(fmt.Printf("O valor do indice é %d e o valor é %d\n", i, v))
	}
}

// go run 3-foundation/5-traversing-arrays/*.go
