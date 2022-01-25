package main

import (
	"fmt"
	"sync"
	"time"
)

const durationTasks = 2
const countProcessing = 3

func main() {
	var wg sync.WaitGroup
	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	in := make(chan int)
	for i := 0; i < countProcessing; i++ {
		wg.Add(1)
		go worker(in, &wg)
	}

	ticker := time.NewTicker(durationTasks * time.Second)

	for i := 0; i < len(jobs); i += countProcessing {
		<-ticker.C
		for j := 0; j < countProcessing && j+i < len(jobs); j++ {
			fmt.Println("send to channel", jobs[j+i])
			in <- jobs[j+i]
		}
	}
	close(in)
	wg.Wait()
	ticker.Stop()
}

func worker(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		result, ok := <-in
		if !ok {
			fmt.Println("close channel")
			return
		}
		fmt.Println("starting in", result)
		time.Sleep(time.Duration(result) * time.Second)
		fmt.Println("finish in", result)
	}
}
