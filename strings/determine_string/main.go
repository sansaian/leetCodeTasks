package main

import "fmt"

func main() {
	fmt.Println(halvesAreAlike("textbook"))
}

func halvesAreAlike(s string) bool {
	var aCount, bCount int
	a, b := s[:len(s)/2], s[len(s)/2:]
	for i := 0; i < len(s)/2; i++ {
		if _, ok := vowels[a[i]]; ok {
			aCount++
		}
		if _, ok := vowels[b[i]]; ok {
			bCount++
		}
	}
	return aCount == bCount
}

var vowels = map[byte]struct{}{
	'a': {},
	'A': {},
	'e': {},
	'E': {},
	'i': {},
	'I': {},
	'o': {},
	'O': {},
	'u': {},
	'U': {},
}
