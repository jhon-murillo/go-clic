package main

import (
	"log"
	"net"
	"net/http"
	"net/url"
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
	
	u, err = url.ParseRequestURI("golang.org")
		if err != nil {
		panic(err)
	}
	log.Printf("err=%+v url=%+v\n", err, u)
	
	
	resp, err := client.Get("https://golang.org/")
		if err != nil {
		panic(err)
	}
	log.Println(resp)
}

func isValidUri(uri string) bool {
  _, err := url.ParseRequestURI(uri)

  return err == nil
}
