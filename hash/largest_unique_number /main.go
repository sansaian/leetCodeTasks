package main

func main() {

}

func largestUniqueNumber(nums []int) int {
	store := make(map[int]int, len(nums))
	for _, v := range nums {
		store[v] = store[v] + 1
	}
	max := -1
	for k, v := range store {
		if v != 1 {
			continue
		}
		if max < k {
			max = k
		}
	}
	return max
}
