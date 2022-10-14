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
	"sync"
	
)

func main() {
	
	if len(os.Args) == 1 {
	    log.Println("Usage: %s URL", filepath.Base(os.Args[0]))
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
	
	var wg sync.WaitGroup
	
	for _, rawUrl := range os.Args[1:] {
	
	    wg.Add(1)	
		
	    go func(rawUrl string) {
	    
	    	defer wg.Done()
		
		u, err := url.ParseRequestURI(rawUrl)
		resp, err := client.Get(u.String())
		
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
                size := len(body)
	        log.Println(u , size)
	    
	    }(rawUrl)
	}
	wg.Wait()
	defer resp.Body.Close()
	
}
