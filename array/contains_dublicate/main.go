package main

func main() {

}

func containsDuplicate(nums []int) bool {
	var x = make(map[int]bool)
	for _, s := range nums {
		if x[s] {
			return true
		}
		x[s] = true
	}
	return false
}
