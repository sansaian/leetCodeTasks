package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 1
	rotate(nums, k)
	fmt.Println(nums)

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	k = 2
	rotate(nums, k)
	fmt.Println(nums)

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	k = 3
	rotate(nums, k)
	fmt.Println(nums)

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	k = 8
	rotate(nums, k)
	fmt.Println(nums)

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	k = 16
	rotate(nums, k)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	if len(nums) < k {
		k = k % len(nums)
	}
	h := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	for i, v := range h {
		nums[i] = v
	}
}
