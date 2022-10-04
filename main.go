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
	
        for i=1, url := range os.Args {
		u, err := url.ParseRequestURI(os.Args[i])
	        if err != nil {
	        panic(err)
	}
        
        
    }
	
	for i, element := range someSlice {
   
}
	
	u, err := url.ParseRequestURI(os.Args[1])
	if err != nil {
	    panic(err)
	}
	
	f, err := os.OpenFile("url.list", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
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

func UrlToLines(url string) ([]string, error) {
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
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return LinesFromReader(resp.Body)
}

func LinesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func newWatcher() watcher {
	w := watcher{}
	go w.watch()
	return w
}

type watcher struct { /* Some resources */
}

func (w watcher) watch() {}

func (w watcher) close() {
	// Close the resources
}
