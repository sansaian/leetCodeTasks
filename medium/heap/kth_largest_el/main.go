package main

import (
	"slices"
)

func main() {

}

func findKthLargest(nums []int, k int) int {
	slices.Sort(nums)
	return nums[len(nums)-k]
}
