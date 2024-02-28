package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{0, 1}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	fmt.Println(missingNumber([]int{1}))

}

func missingNumber(nums []int) int {
	hm := make(map[int]struct{})
	for _, element := range nums {
		hm[element] = struct{}{}
	}
	for i := 0; i <= len(nums); i++ {
		if _, ok := hm[i]; !ok {
			return i
		}
	}
	return -1
}
