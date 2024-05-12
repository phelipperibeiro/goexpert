package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Unidade     string `json:"unidade"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
}

func main() {
	fmt.Println("Server is running on port 8080")
	http.HandleFunc("/", SearchCepHandler)
	http.ListenAndServe(":8080", nil)

}

func SearchCepHandler(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		response.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := request.URL.Query().Get("cep")
	if cepParam == "" {
		response.Write([]byte(`{"message": "cep is required"}`))
		fmt.Println("cep is required")
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	cep, error := searchCep(cepParam)
	if error != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	// result, err := json.Marshal(cep)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(result)
	json.NewEncoder(response).Encode(cep)
}

func searchCep(cep string) (*ViaCEP, error) {
	resp, error := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}
	var c ViaCEP
	error = json.Unmarshal(body, &c)
	if error != nil {
		return nil, error
	}
	return &c, nil
}

// go run 4-important-packages/5-api/api-http/main.go
