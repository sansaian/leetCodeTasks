package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
}

func subarraySum(nums []int, k int) int {
	count := 0
	prefixSum := make([]int, 0, len(nums))
	sumCount := map[int]int{0: 1}
	tmp := 0
	// Calculate prefix sum
	for _, v := range nums {
		tmp += v
		prefixSum = append(prefixSum, tmp)
	}
	// Check for subarrays with sum k using prefix sums
	for _, curr := range prefixSum {
		count += sumCount[curr-k]
		sumCount[curr] += 1
	}

	return count
}
