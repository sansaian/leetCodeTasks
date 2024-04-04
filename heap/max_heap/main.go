package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &MaxHeap{10, 5, 20, 7, 15}

	// Initialize heap
	heap.Init(h)

	// Insert into heap
	heap.Push(h, 25)

	// Extract maximum element
	max := heap.Pop(h).(int)
	fmt.Println("Max:", max)

	max = heap.Pop(h).(int)
	fmt.Println("Max:", max)

	// Insert into heap
	heap.Push(h, 30)

	max = (*h)[0]
	fmt.Println("Max:", max)
}
