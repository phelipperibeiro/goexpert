package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// dd é uma função que imprime informações de debug e interrompe a execução do programa
func dd(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Erro ao serializar dados:", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData))
	os.Exit(0)
}

func main() {
	// Exemplo de uso do dd
	exampleData := map[string]interface{}{
		"name":  "John Doe",
		"age":   30,
		"email": "john@example.com",
	}

	// Chamada ao dd para imprimir e interromper a execução
	dd(exampleData)

	// Este código nunca será alcançado devido à chamada ao dd
	fmt.Println("Este código nunca será alcançado")
}

// go run 2-dd/main.go
