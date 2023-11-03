package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumSum([]int{368, 369, 307, 304, 384, 138, 90, 279, 35, 396, 114, 328, 251, 364, 300, 191, 438, 467, 183}))
	fmt.Println(maximumSum([]int{18, 43, 36, 13, 7}))

	fmt.Println(maximumSum([]int{10, 12, 19, 14}))
}

func maximumSum(nums []int) int {
	// Map to store the maximum two numbers for each digit sum
	store := make(map[int][]int)
	result := -1

	for _, n := range nums {
		digitSum := calculateSum(n)

		// Check if there are already values for this digit sum
		values, exists := store[digitSum]
		if !exists {
			// If not, just add the number to the store
			store[digitSum] = []int{n}
		} else {
			// Otherwise, update the existing values
			values = append(values, n) // Add the current number to the values

			// Sort the values to have the largest numbers at the beginning
			sort.Sort(sort.Reverse(sort.IntSlice(values)))

			// Keep only the two largest numbers if there are more than two
			if len(values) > 2 {
				values = values[:2]
			}

			// Save back to the store
			store[digitSum] = values

			// Also, update the result if we have at least two numbers
			if len(values) > 1 {
				result = max(result, values[0]+values[1])
			}
		}
	}

	return result
}

// Helper function to calculate the sum of digits of a number
func calculateSum(n int) int {
	sum := 0
	for n != 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// Helper function to find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
