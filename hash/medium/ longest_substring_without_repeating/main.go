package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	store := make(map[uint8]int)
	var left, right, result int

	for right < len(s) {
		var r = s[right]
		store[r] += 1
		for store[r] > 1 {
			l := s[left]
			store[l] -= 1
			left += 1
		}
		result = max(result, right-left+1)

		right += 1
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
