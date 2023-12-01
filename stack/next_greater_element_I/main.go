package main

func main() {

}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	mapstore := make(map[int]int, len(nums1))
	var stack []int
	for _, v := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < v {
			mapstore[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
	}
	for i, v := range nums1 {
		if mapv, ok := mapstore[v]; ok {
			nums1[i] = mapv
			continue
		}
		nums1[i] = -1
	}
	return nums1
}
