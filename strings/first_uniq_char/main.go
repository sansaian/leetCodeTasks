package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
	fmt.Println(firstUniqChar("aabb"))
}

func firstUniqChar(s string) int {
	heap := make(map[int32]int, len(s))
	for _, value := range s {
		heap[value]++
	}
	for i, val := range s {
		if heap[val] == 1 {
			return i
		}
	}
	return -1
}
