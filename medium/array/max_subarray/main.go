package main

import "math"

func main() {

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
