package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	var profit int
	low, profit := 10000, 0
	for _, current := range prices {
		if current < low {
			low = current
			continue
		}
		if profit < (current - low) {
			profit = current - low
		}
	}
	return profit
}
