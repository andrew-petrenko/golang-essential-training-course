package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			contentType(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func contentType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s\n", url, ctype)
}
