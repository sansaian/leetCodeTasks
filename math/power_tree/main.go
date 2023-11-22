package main

import "fmt"

func main() {
	fmt.Println(isPowerOfThree(1))
	fmt.Println(isPowerOfThree(15))
	fmt.Println(isPowerOfThree(27))
	fmt.Println(isPowerOfThree(45))
}

func isPowerOfThree(n int) bool {
	// 1162261467 is 3^19,  3^20 is bigger than int
	return n > 0 && 1162261467%n == 0
}
