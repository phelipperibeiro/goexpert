package main

import "fmt"

type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func (conta Conta) simular(valor int) int {
	conta.saldo += valor
	fmt.Println("Saldo simulado:", conta.saldo)
	return conta.saldo
}

func (conta *Conta) adicionar(valor int) int {
	conta.saldo += valor
	fmt.Println("Saldo creditado:", conta.saldo)
	return conta.saldo
}

func main() {

	conta1 := NewConta()
	conta1.simular(200)
	fmt.Println("Saldo da conta nº1:", conta1.saldo)
	conta1.adicionar(100)
	fmt.Println("Saldo da conta nº1:", conta1.saldo)

	conta2 := NewConta()
	conta2.simular(300)
	fmt.Println("Saldo da conta nº2:", conta2.saldo)
	conta2.adicionar(200)
	fmt.Println("Saldo da conta nº2:", conta2.saldo)
}

// go run 3-foundation/12.2-pointers-use-case/*.go
