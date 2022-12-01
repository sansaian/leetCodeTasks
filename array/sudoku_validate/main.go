package main

import "strconv"

func main() {

}

func isValidSudoku(board [][]byte) bool {
	store := make(map[string]struct{})

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if string(board[row][col]) == "." {
				continue
			}

			if _, ok := store[string(board[row][col])+"found_row"+strconv.Itoa(row)]; ok {
				return false
			}
			store[string(board[row][col])+"found_row"+strconv.Itoa(row)] = struct{}{}

			if _, ok := store[string(board[row][col])+"found_column"+strconv.Itoa(col)]; ok {
				return false
			}
			store[string(board[row][col])+"found_column"+strconv.Itoa(col)] = struct{}{}

			if _, ok := store[string(board[row][col])+"found_box"+strconv.Itoa(row/3)+"-"+strconv.Itoa(col/3)]; ok {
				return false
			}
			store[string(board[row][col])+"found_box"+strconv.Itoa(row/3)+"-"+strconv.Itoa(col/3)] = struct{}{}

		}
	}
	return true
}
