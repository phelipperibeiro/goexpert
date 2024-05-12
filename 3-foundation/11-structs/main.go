package main

import "fmt"

/*
 * #############################################
 * Pessoa Interface
 * #############################################
 */
type Pessoa interface {
	Desativar()
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

/*
 * #############################################
 * Cliente
 * #############################################
 */
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

/*
 * #############################################
 * Empresa
 * #############################################
 */

type Empresa struct {
	Nome  string
	Ativo bool
	Endereco
}

func (e Empresa) Desativar() {
	fmt.Printf("A empresa %s foi desativada\n", e.Nome)
}

/*
 * #############################################
 * Endereco
 * #############################################
 */
type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

func main() {

	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
		Endereco: Endereco{
			Logradouro: "Rua 1",
		},
	}

	empresa := Empresa{
		Nome:  "Perfumaria XPTO",
		Ativo: true,
		Endereco: Endereco{
			Logradouro: "Rua 2",
		},
	}

	Desativacao(wesley)
	Desativacao(empresa)
}

// go run 3-foundation/11-structs/*.go
