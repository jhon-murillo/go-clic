package main

import (
	"log"
	"github.com/hashicorp/go-retryablehttp"
	"net"
	"net/http"
	"time"
)

func main() {
	
	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: time.Second,
			}).DialContext,
			TLSHandshakeTimeout:   time.Second,
			ResponseHeaderTimeout: time.Second,
		},
	}
	
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
