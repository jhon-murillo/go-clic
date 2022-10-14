package main

import (
	
	"os"
	"path/filepath"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
	"io"
	
)

func main() {
	
	if len(os.Args) == 1 {
	    log.Println("Usage: %s URL", filepath.Base(os.Args[0]))
	}
	
	var u *url.URL
	var err error
	var resp *http.Response
	var size []byte
	
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
	
	for _, rawUrl := range os.Args[1:] {
	    u, err = url.ParseRequestURI(rawUrl)
            resp, err = client.Get(u.String())
	    if err != nil {
	        panic(err)
	    }
	    if resp.StatusCode != http.StatusOK {
	        log.Println (u.String(), "Status code: ", resp.StatusCode)
	    }
	    body, err := io.ReadAll(resp.Body) 
	    if err != nil {
	        panic(err)
	    }
            size = len(body)   	
	}
	defer resp.Body.Close()
}
