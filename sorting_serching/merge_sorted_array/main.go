package main

import "fmt"

func merge(nums1 []int, nums2 []int) []int {
	result := make([]int, 0, len(nums1)+len(nums2))
	var j, k int

	for i := 0; i <= len(nums1)+len(nums2); i++ {
		if j < len(nums1) && k < len(nums2) {
			if nums1[j] <= nums2[k] {
				result = append(result, nums1[j])
				j++
				continue
			} else {
				result = append(result, nums2[k])
				k++
				continue
			}
		}
		if j < len(nums1) {
			result = append(result, nums1[j])
			j++
			continue
		} else if k < len(nums2) {
			result = append(result, nums2[k])
			k++
		}

	}
	return result
}

func mergeLeetCode(nums1 []int, m int, nums2 []int, n int) {

	copyNyms1 := make([]int, m)
	copy(copyNyms1, nums1[:m])
	var j, k int

	for i := 0; i < m+n; i++ {
		if j < m && k < n {
			if copyNyms1[j] <= nums2[k] {
				nums1[i] = copyNyms1[j]
				j++
				continue
			} else {
				nums1[i] = nums2[k]
				k++
				continue
			}
		}
		if j < m {
			nums1[i] = copyNyms1[j]
			j++
		}
		if k < n {
			nums1[i] = nums2[k]
			k++
		}
	}
}

func main() {
	fmt.Println(merge([]int{3, 5, 8, 9}, []int{2, 5, 6, 10}))
	fmt.Println(merge([]int{1, 2, 3, 6, 8}, []int{2, 5, 6}))

	mergeLeetCode([]int{3, 5, 8, 9, 0, 0, 0, 0}, 4, []int{2, 5, 6, 10}, 4)
	mergeLeetCode([]int{1, 2, 3, 6, 8, 0, 0, 0}, 4, []int{2, 5, 6}, 3)

}
