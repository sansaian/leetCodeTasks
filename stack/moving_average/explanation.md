# Intuition
Initially, the problem seems to involve maintaining a running total of the latest `n` numbers in a sequence and calculating their average efficiently. The challenge lies in efficiently updating the average as new numbers are added, especially when the sequence grows large. A sliding window approach using a prefix sum array seems like a natural solution to avoid recalculating the sum of the window each time.

# Approach
The solution employs a data structure, `MovingAverage`, which maintains a running total of numbers and their counts to compute the moving average. The key is to use a prefix sum array (`sumPrfx`) which keeps cumulative sums of the numbers. When a new number is added, the cumulative sum is updated accordingly. For the moving average calculation, we subtract the sum just before the window starts from the current sum, thus obtaining the sum within the window. We then divide this sum by the window's size to get the average.

# Complexity
- Time complexity:
  The time complexity of adding a new value (the `Next` method) is $$O(1)$$ as it involves basic arithmetic operations and appending to a slice in Go, which is generally an $$O(1)$$ operation, thanks to dynamic resizing.

- Space complexity:
  The space complexity is $$O(N)$$, where `N` is the number of elements added to the `MovingAverage`. This is due to the storage requirements of the `sumPrfx` slice, which grows with each new element.

# Code
```go
type MovingAverage struct {
    sumPrfx []int
    windowSize  int
}

func Constructor(size int) MovingAverage {
    return MovingAverage{
        sumPrfx: []int{0},
        windowSize: size,
    }
}

func (this *MovingAverage) Next(val int) float64 {
    this.sumPrfx = append(this.sumPrfx,this.sumPrfx[len(this.sumPrfx)-1]+ val)
    var start int
    if len(this.sumPrfx) > this.windowSize {
        start = len(this.sumPrfx)-1 - this.windowSize
    }
    sum := this.sumPrfx[len(this.sumPrfx)-1] - this.sumPrfx[start]
    return float64(sum)/float64(len(this.sumPrfx)-start-1)
}
```