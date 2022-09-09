package main

import "fmt"

func main() {
	fmt.Println(matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4))
	fmt.Println(matrixReshape([][]int{{1, 2}}, 1, 1))
}

func matrixReshape(mat [][]int, r int, c int) [][]int {
	var store []int
	var result [][]int
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			store = append(store, mat[i][j])
		}
	}
	var pointer int
	for row := 0; row < r; row++ {
		if pointer >= len(store) {
			return mat
		}
		result = append(result, []int{})
		for column := 0; column < c; column++ {
			result[row] = append(result[row], store[pointer])
			pointer++
		}
	}
	if pointer != len(store) {
		return mat
	}
	return result
}
