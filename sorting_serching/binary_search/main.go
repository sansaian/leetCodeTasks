package main

import "fmt"

func main() {
	fmt.Println(Search([]int{1, 2, 3, 4, 5}, 4))
	fmt.Println(Search([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(Search([]int{1, 2}, 1))
}

func Search(arr []int, value int) int {
	return searchBinary(arr, value, 0, len(arr)-1)
}

func searchBinary(arr []int, value int, low int, high int) int {
	if low >= high {
		return -1
	}
	mid := low + (high-low)/2

	if arr[mid] == value {
		return mid
	}
	if arr[mid] < value {
		return searchBinary(arr, value, mid+1, high)
	}
	if arr[mid] > value {
		return searchBinary(arr, value, low, mid-1)
	}
	return -1
}
