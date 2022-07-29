package main

import (
	"fmt"
	"net/http"
)

var (
	urls        = []string{"https://google.com", "https://amazon.com"}
	countWorker = 2
)
var errValue = "-1"

func main() {
	in := make(chan string)
	out := make(chan string)

	for i := 0; i < countWorker; i++ {
		go worker(in, out)
	}
	for _, url := range urls {
		in <- url
	}
	for range urls {
		status := <-out
		fmt.Println(status)
	}
	close(in)
}

func worker(in <-chan string, out chan<- string) {
	for {
		url, ok := <-in
		if !ok {
			return
		}
		status, err := getStatus(url)
		if err != nil {
			fmt.Println(fmt.Errorf("url %s, %w", url, err))
			out <- errValue
			return
		}
		out <- status
	}
}

func getStatus(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	return resp.Status, nil
}
