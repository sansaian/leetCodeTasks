package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(minStoneSum([]int{5, 4, 9}, 2))
}

func minStoneSum(piles []int, k int) int {
	var result int
	h := &MaxHeap{}
	heap.Init(h)
	for _, num := range piles {
		val := num
		result += val
		heap.Push(h, val)
	}
	for i := 0; i < k; i++ {
		largest := heap.Pop(h).(int)
		reduction := largest / 2
		result -= reduction
		heap.Push(h, largest-reduction)
	}
	return result
}

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
