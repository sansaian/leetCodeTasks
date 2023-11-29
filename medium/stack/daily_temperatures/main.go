package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0, len(temperatures))
	ans := make([]int, len(temperatures), len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[j] = i - j
		}
		stack = append(stack, i)
	}
	return ans
}
