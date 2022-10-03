package main

import (
	
	"os"
	"path/filepath"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"

)

func main() {
	if len(os.Args) != 2 {
	    log.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
	}
	
	u, err := url.ParseRequestURI(os.Args[1])
	if err != nil {
	    panic(err)
	}
	log.Printf("err=%+v url=%+v\n", err, u)
	
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
	
	resp, err := client.Get(u.String())
	if err != nil {
	    panic(err)
	}
	log.Println(resp)
	
	if resp.StatusCode != http.StatusOK {
	}
	log.Printf("Status code:", resp.StatusCode)
}
