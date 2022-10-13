package main

import (
	
	"os"
	"path/filepath"
	"log"
	"sync"
	"net"
	"net/http"
	"net/url"
	"time"
	"io"
	"sort"
	
)

func main() {
	
	if len(os.Args) == 1 {
	    log.Println("Usage: %s URL", filepath.Base(os.Args[0]))
	}
	
	n := len(os.Args[1:])
	keys := make([]string, 0, n)
	body := make([]byte, 0, n)
	m := make(map[string]int, n)
	
	var u *url.URL
	var err error
	var resp *http.Response
	var wg sync.WaitGroup
        
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
	
	    wg.Add(1)
	    
	    go func(val string) {
		    
		    defer wg.Done()
		    
		    u, err = url.ParseRequestURI(val)
		    resp, err = client.Get(u.String())
		
		    if err != nil {
	    		    panic(err)
		    }
		
		    if resp.StatusCode != http.StatusOK {
			log.Println (u.String(), "Status code: ", resp.StatusCode)
		    }
		    
		    body, err = io.ReadAll(resp.Body) 
		    
	    	    if err != nil {
	    	        panic(err)
	    	    }
		    
		    m[val] = len(body)
		    
                    wg.Done()    
				       
	    }(rawUrl)
	    
	    keys = append(keys, rawUrl)		
	}
	
	wg.Wait()

	sort.SliceStable(keys, func(i, j int) bool{
        	return m[keys[i]] < m[keys[j]]
    	})
	
	for _, k := range keys{
        	log.Println(k, m[k])
    	}
	
	defer resp.Body.Close()
}
