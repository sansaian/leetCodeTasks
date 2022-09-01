package main

import "fmt"

func main() {
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("fffbfg", "effjfggbffjdgbjjhhdegh"))
}

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	store := make(map[rune]int, 26)
	for _, sym := range magazine {
		store[sym]++
	}
	for _, sym := range ransomNote {
		value, ok := store[sym]
		if !ok || value == 0 {
			return false
		}
		store[sym]--
	}
	return true
}
