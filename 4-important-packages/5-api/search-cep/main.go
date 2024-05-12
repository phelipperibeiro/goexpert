package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	for _, cep := range os.Args[1:] {

		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		}
		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
		}
		file, err := os.Create("cidade.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
		}
		defer file.Close()
		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", data.Cep, data.Localidade, data.Uf))
		fmt.Println("Arquivo criado com sucesso!")
		fmt.Println("-------------------------------------------------")
		fmt.Println("Cidade: ", data.Localidade)
		fmt.Println("UF: ", data.Uf)
		fmt.Println("CEP: ", data.Cep)
		fmt.Println("Bairro: ", data.Bairro)
		fmt.Println("Logradouro: ", data.Logradouro)
		fmt.Println("Complemento: ", data.Complemento)
		fmt.Println("IBGE: ", data.Ibge)
		fmt.Println("GIA: ", data.Gia)
		fmt.Println("DDD: ", data.Ddd)
		fmt.Println("SIAFI: ", data.Siafi)
		fmt.Println("-------------------------------------------------")
	}
}

// go run 4-important-packages/5-busca-cep/main.go 05729120 04942000
