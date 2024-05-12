package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	client := http.Client{Timeout: time.Second}
	resp, err := client.Get("http://google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
