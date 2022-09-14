package main

import "fmt"

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	//fmt.Println(searchMatrix([][]int{{1}}, 1))

}

func searchMatrix(matrix [][]int, target int) bool {
	store := make([]int, 0, len(matrix)*len(matrix[0]))
	for i := range matrix {
		for j := range matrix[i] {
			store = append(store, matrix[i][j])
		}
	}
	return searchBinary(store, target, 0, len(store))
}

func searchBinary(arr []int, value int, low int, high int) bool {
	if low >= high {
		return false
	}
	mid := low + (high-low)/2

	if arr[mid] == value {
		return true
	}
	if arr[mid] < value {
		return searchBinary(arr, value, mid+1, high)
	}
	if arr[mid] > value {
		return searchBinary(arr, value, low, mid)
	}
	return false
}
