package main

import "fmt"

func main() {
	fmt.Println(repeatedCharacter("abccbaacz"))
}

func repeatedCharacter(s string) byte {
	set := make(map[byte]struct{})
	for i := 0; i < len(s); i++ {
		if _, ok := set[s[i]]; ok {
			return s[i]
		}
		set[s[i]] = struct{}{}
	}
	return ' '
}
