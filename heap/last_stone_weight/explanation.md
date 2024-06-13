# Problem Statement
The `KthLargest` class is designed to manage a stream of numbers and allow for the efficient retrieval of the k-th largest element in the stream.

## Example Usage
- **Input**: Initialize with `k = 3` and `nums = [4, 5, 8, 2]`, operations: `add(3)`, `add(5)`, `add(10)`, `add(9)`, `add(4)`
- **Output**: After each add operation, it returns the k-th largest element from the stream.

# Approach
The approach involves using a min-heap to keep track of the top `k` elements. The smallest element in the heap is the k-th largest element in the data stream:
1. **Constructor**: Initialize the min-heap and fill it with the first `k` elements from the input list. For any additional elements, add them to the heap only if they are larger than the smallest element (i.e., the root of the heap), and remove the smallest element.
2. **Add method**: Insert a new value into the heap. If the heap size exceeds `k`, remove the smallest element, ensuring the heap always contains exactly `k` largest elements. The root of the heap represents the k-th largest element.

### Steps
1. Initialize a `MinHeap` in the constructor and add elements up to the size `k`.
2. For each new element added, compare it with the root. If it is larger, pop the root and push the new element.
3. Always return the root of the heap after adding an element, as it represents the k-th largest element.

## Complexity Analysis
- **Time Complexity**:
    - Constructor: O(n log k), where `n` is the number of initial elements and `k` is the heap size.
    - Add method: O(log k), due to the heap operations (push and pop).
- **Space Complexity**: O(k), to store the heap of size `k`.

# Code
```go
package main

import (
	"container/heap"
	"fmt"
)

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

func main() {
	myheap := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(myheap.Add(3))  // Should return the 3rd largest element
	fmt.Println(myheap.Add(5))
	fmt.Println(myheap.Add(10))
	fmt.Println(myheap.Add(9))
	fmt.Println(myheap.Add(4))
}
