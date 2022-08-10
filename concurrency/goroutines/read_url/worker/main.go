package main

import (
	"fmt"
	"time"
)

const n = 2

func readURL(in <-chan string, out chan<- string) {
	//var result strings
	for {
		result, ok := <-in
		fmt.Printf("get url %s \n", result)
		if ok == false {
			fmt.Println("close goroutine")
			return
		}
		time.Sleep(3 * time.Second)
		fmt.Printf("read url %s \n", result)
		out <- result
	}

}

func main() {
	job := []string{"1", "2", "3", "4"}
	in := make(chan string)
	out := make(chan string, len(job))
	//create workers
	for i := 0; i < n; i++ {
		go readURL(in, out)
	}
	// send job
	for _, v := range job {
		in <- v
	}

	//get result
	for j := 1; j <= len(job); j++ {
		<-out
	}
	close(in)
}
