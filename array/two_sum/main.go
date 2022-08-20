package main

func main() {

}

func twoSum(nums []int, target int) []int {
	mapValue := make(map[int]int)
	for index, element := range nums {
		temp := target - element
		if val, ok := mapValue[element]; ok {
			return []int{val, index}
		}
		mapValue[temp] = index
	}
	return []int{0, 0}
}
