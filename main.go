package main

import (
	"net/http"
	"bufio"
	"io"
	"os"
)

func main() {
	listing()
	file, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	_, _ = countEmptyLines(file)
}

func countEmptyLines(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		// ...
	}
	return 0, nil
}

func listing() error {
	var client *http.Client
	var err error
	if tracing {
		client, err = createClientWithTracing()
	} else {
		client, err = createDefaultClient()
	}
	if err != nil {
		return err
	}

	_ = client
	return nil
}

var tracing bool

func createClientWithTracing() (*http.Client, error) {
	return nil, nil
}

func createDefaultClient() (*http.Client, error) {
	return nil, nil
}
