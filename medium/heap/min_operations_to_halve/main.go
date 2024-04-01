package main

import (
	"container/heap"
	"fmt"
)

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	fmt.Println(halveArray(stones))

}
func halveArray(nums []int) int {
	var target float64
	h := &MaxHeap{}
	heap.Init(h)
	for _, num := range nums {
		val := float64(num)
		target += val
		heap.Push(h, val)
	}
	target /= 2

	var sum float64
	ops := 0
	for sum < target {
		ops++
		largest := heap.Pop(h).(float64)
		half := largest / 2
		sum += half
		heap.Push(h, half)
	}
	return ops
}

type MaxHeap []float64

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(float64))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
