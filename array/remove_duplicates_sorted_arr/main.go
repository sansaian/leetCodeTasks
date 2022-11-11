package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func removeDuplicates(nums []int) int {
	var s int
	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[s] {
			continue
		}
		s++
		nums[s] = nums[i]
	}
	return s + 1
}
