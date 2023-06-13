package main

import "fmt"

func main() {
	fmt.Println(runningSum([]int{1, 2, 3, 4}))
}

func runningSum(nums []int) []int {
	prefix := make([]int, len(nums)+1)
	for i := 1; i < len(prefix); i++ {
		prefix[i] = prefix[i-1] + nums[i-1]
	}
	return prefix[1:]
}
