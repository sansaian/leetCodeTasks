[# Problem Statement
Given an array of coordinate points and an integer `k`, return the `k` closest points to the origin (0, 0) based on their Euclidean distance.
]()
## Example
- **Input**: points = [[1,3], [3,4], [2,-1]], k = 2
- **Output**: [[1,3], [2,-1]]
- **Explanation**: The distances of points from the origin are [sqrt(10), 5, sqrt(5)], respectively. The two closest points are [1,3] and [2,-1].

# Approach
To solve this problem efficiently, a max heap is used to keep track of the `k` closest points by their Euclidean distances. By maintaining the size of the heap to `k`, we ensure that we always discard the farthest point when the heap size exceeds `k`.

### Steps
1. Initialize a max heap that will store points with their distances from the origin.
2. Iterate over each point in the given list:
    - Calculate the Euclidean distance of the point from the origin.
    - Push the point with its distance into the heap.
    - If the heap size exceeds `k`, pop the heap to remove the point farthest from the origin.
3. After processing all points, the heap contains the `k` closest points.
4. Extract these points from the heap and add them to the result list in the correct order.

## Complexity Analysis
- **Time Complexity**: O(n log k), where `n` is the number of points. Each insertion into a heap of size `k` takes O(log k) time.
- **Space Complexity**: O(k) for storing the heap.

# Code
```go
import (
	"container/heap"
	"math"
)

type Point struct {
	distance float64
	x        int
	y        int
}

type PointHeap []Point

func (h PointHeap) Len() int { return len(h) }
func (h PointHeap) Less(i, j int) bool {
	return h[i].distance > h[j].distance // Max-heap based on distance
}

func (h PointHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PointHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *PointHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	h := &PointHeap{}
	heap.Init(h)
	for _, point := range points {
		distance := math.Sqrt(float64(point[0]*point[0] + point[1]*point[1]))
		heap.Push(h, Point{distance: distance, x: point[0], y: point[1]})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	result := make([][]int, k)
	for i := 0; i < k; i++ {
		tmp := heap.Pop(h).(Point)
		result[i] = []int{tmp.x, tmp.y}
	}
	return result
}
