package main

import "fmt"

func main() {
	fmt.Println(reversePrefix("abcdefd", 'd'))
	fmt.Println(reversePrefix("xyxzxe", 'z'))
	fmt.Println(reversePrefix("abcd", 'z'))
}

func reversePrefix(word string, ch byte) string {
	result := make([]byte, 0, len(word))
	r := 0
	for i := range word {
		if word[i] == ch {
			break
		}
		r++
	}
	if r == len(word) {
		return word
	}
	for j := r; j >= 0; j-- {
		result = append(result, word[j])
	}
	result = append(result, word[r+1:]...)
	return string(result)
}
