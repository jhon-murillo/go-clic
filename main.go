package main

import (
	"log"
        "client"
	"github.com/hashicorp/go-retryablehttp"
)

func main() {
	
	resp, err := client.Get("https://golang.org/")
		if err != nil {
		panic(err)
	}
	log.Println(resp)
	
	resp, err := retryablehttp.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	log.Println(resp)
}
