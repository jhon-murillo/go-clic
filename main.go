package main

import (
	"log"

	"github.com/hashicorp/go-retryablehttp"
)

func main() {
	resp, err := retryablehttp.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	log.Println(resp)
}
