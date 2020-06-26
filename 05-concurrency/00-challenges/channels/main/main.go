package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	ch := make(chan string)

	for _, url := range urls {
		go contentType(url, ch)
	}

	for range urls {
		out := <-ch
		fmt.Println(out)
	}
}

func contentType(url string, out chan string) {
	resp, err := http.Get(url)
	if err != nil {
		out <- fmt.Sprintf("%s -> error: %s", url, err)
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("Content-Type")
	out <- fmt.Sprintf("%s -> %s", url, ctype)
}
