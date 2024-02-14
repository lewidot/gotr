package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	// send request to the api
	baseURL := "https://the-one-api.dev/v2/"

	resp, err := http.Get(baseURL + "book")
	if err != nil {
		log.Fatalln(err)
	}

	// parse response body
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	fmt.Println(sb)
}
