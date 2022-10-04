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
        var u *URL
	var err error
	
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
	
	if len(os.Args) == 1 {
	    log.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
	}
	
        for i, _ := range os.Args {
		u, err = url.ParseRequestURI(os.Args[i])
	        if err != nil {
	            panic(err)
		}
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
