package main

import "fmt"

func main() {
	fmt.Println(minStartValue([]int{-3, 2, -3, 4, 2}))
}

func minStartValue(nums []int) int {
	sum, minVal := 0, 0
	for _, v := range nums {
		sum += v
		minVal = min(minVal, sum)
	}
	return -minVal + 1
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
