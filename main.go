package main

import (
	
	"os"
	"path/filepath"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
	"io/ioutil"

)

func main() {
	if len(os.Args) != 2 {
	    log.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
	}
	
	u, err := url.ParseRequestURI(os.Args[1])
	if err != nil {
	    panic(err)
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
	    log.Printf(u, "Status code:", resp.StatusCode)
	}
	
	body, error := ioutil.ReadAll(resp.Body)
	if err != nil {
	    panic(err)
	}
	defer resp.Body.Close()
	log.Printf("Body Size:", body)
	
}
