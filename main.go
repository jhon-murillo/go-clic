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
	
	if len(os.Args) == 1 {
	    log.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
	}
	
        for i, url := range os.Args {
		u, err := url.ParseRequestURI(os.Args[i])
	        if err != nil {
	        panic(err)
		}
	}
	
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
	
	if resp.StatusCode != http.StatusOK {
		log.Printf(u.String(), "Status code:", resp.StatusCode)
	}
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	    panic(err)
	}
	log.Printf("Body Size:", string(body))
	
	defer resp.Body.Close()
}
