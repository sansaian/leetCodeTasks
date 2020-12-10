package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse(123))        // 321
	fmt.Println(reverse(-123))       // -321
	fmt.Println(reverse(120))        // 21
	fmt.Println(reverse(0))          // 0
	fmt.Println(reverse(1534236469)) // 0
}

func reverse(x int) int {
	multiplier := 1
	if x < 0 {
		multiplier = -1
		x = x * multiplier
	}
	result := 0
	for x > 0 {
		tmp := x % 10
		x = x / 10
		result = result*10 + tmp
	}
	if result >2147483648 {
		return 0
	}
	return result*multiplier
}
