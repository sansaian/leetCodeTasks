package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4, 1}))
}

func maxSubArray(nums []int) int {
	var max, currSum = math.MinInt, 0
	for i := 0; i < len(nums); i++ {
		if (currSum + nums[i]) < nums[i] {
			currSum = nums[i]
		} else {
			currSum += nums[i]
		}
		if max < currSum {
			max = currSum
		}
	}
	return max
}
