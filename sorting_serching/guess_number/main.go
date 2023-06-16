package main

import "fmt"

const goal = 6

func main() {
	fmt.Println(guessNumber(1))
}

func guessNumber(n int) int {
	return binarySearch(0, n)
}

func binarySearch(start, finish int) int {
	if start > finish {
		return -1
	}
	dot := (finish-start)/2 + start

	switch guess(dot) {
	case -1:
		return binarySearch(start, dot-1)
	case 1:
		return binarySearch(dot+1, finish)
	default:
		return dot
	}
}

func guess(dot int) int {
	if dot < goal {
		return 1
	}
	if dot > goal {
		return -1
	}
	return 0
}
