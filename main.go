package main

import (
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
	"os"
	"path/filepath"
)

func main() {
	var URL string
	
	if len(os.Args) != 2 {
	    fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
	    URL = "google.com" 	
	}
	
	URL = os.Args[1]
	u, err := url.ParseRequestURI(URL)
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
