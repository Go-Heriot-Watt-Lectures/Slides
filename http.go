package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting http request")
	response, err := http.Get("http://google.com")
	if err != nil {
		panic("Problem during get request")
	}
	if (response.StatusCode == http.StatusOK) {
		log.Println("Request succeeded!")
	}
}
