package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}

func pivotIndex(nums []int) int {
	var sum, leftSum int
	for _, value := range nums {
		sum += value
	}
	for i := range nums {
		if leftSum == sum-leftSum-nums[i] {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}
