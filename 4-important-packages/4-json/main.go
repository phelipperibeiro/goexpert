package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

/**
 * json.Marshal returns a byte slice with the json representation of the struct
 */
func example1() {
	println("Example 1")
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string(res))
}

/**
 * json.NewEncoder writes the json representation of the struct to the writer
 */
func example2() {
	println("Example 2")
	conta := Conta{Numero: 1, Saldo: 100}
	err := json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}
}

/**
 * json.Unmarshal parses the json representation of the struct and stores it in the struct
 */
func example3() {
	println("Example 3")

	jsonPuro := []byte(`{"n":2,"s":200}`)
	var contaX Conta
	err := json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		println(err)
	}
	println(contaX.Saldo)
}

func main() {
	example1()
	example2()
	example3()
}

// go run 4-important-packages/4-json/main.go
