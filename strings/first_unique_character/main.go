package main

import "fmt"

func main() {

	fmt.Println(firstUniqChar("sseretv"))
}

func firstUniqChar(s string) int {
	heapmap := make(map[int32]int, len(s))
	for _, value := range s {
		heapmap[value]++
	}
	for i, val := range s {
		if heapmap[val] == 1 {
			return i
		}
	}
	return -1
}
