package main

import "fmt"

func main() {
	fmt.Println(addDigits(38))
}

func addDigits(num int) int {
	var result int
	for {
		for num > 0 {
			result += num % 10
			num /= 10
		}
		num, result = result, 0
		fmt.Println(num, " ", result)
		if num < 10 {
			break
		}
	}
	return num
}
