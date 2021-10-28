package main

import (
	"fmt"
)

func main() {

	matrix1 := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	matrix2 := [][]int{{1, 2}, {3, 4}}
	matrix3 := [][]int{{1}}
	rotate(matrix1)
	fmt.Println(matrix1)
	rotate(matrix2)
	fmt.Println(matrix2)
	rotate(matrix3)
	fmt.Println(matrix3)
}

func rotate(matrix [][]int) {
	lenght := len(matrix[0])
	duplicate := make([][]int, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]int, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}

	for i := 0; i < lenght; i++ {
		for j := 0; j < lenght; j++ {
			if i == lenght && j == lenght {
				continue
			}
			matrix[j][(lenght-1)-i] = duplicate[i][j]
		}
	}
	matrix = duplicate
}
