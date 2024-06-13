package main

import (
	"container/heap"
	"fmt"
)

func main() {
	myheap := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(myheap.Add(3))
	fmt.Println(myheap.Add(5))
	fmt.Println(myheap.Add(10))
	fmt.Println(myheap.Add(9))
	fmt.Println(myheap.Add(4))
}

type KthLargest struct {
	MaxHeap *MinHeap
	Limit   int
}

func Constructor(k int, nums []int) KthLargest {

	h := &MinHeap{}
	heap.Init(h)
	for _, el := range nums {
		heap.Push(h, el)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return KthLargest{
		MaxHeap: h,
		Limit:   k,
	}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.MaxHeap, val)
	if this.MaxHeap.Len() > this.Limit {
		heap.Pop(this.MaxHeap)
	}
	return this.MaxHeap.Min().(int)

}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Min() (v any) {
	return (*h)[0]
}

func (h *MinHeap) Pop() (v any) {
	v, *h = (*h)[len(*h)-1], (*h)[:len(*h)-1]
	return
}
