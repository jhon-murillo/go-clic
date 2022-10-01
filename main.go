package main

import (
    "github.com/hashicorp/go-retryablehttp"
)

func main() {
    resp, err := retryablehttp.Get("www.google.com")
    if err != nil {
    panic(err)
    }
}
