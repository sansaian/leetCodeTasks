package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("str=asdfgha result=", isUniqueChar("asdfgha"))
	fmt.Println("str=asddfgh result= ", isUniqueChar("asddfgh"))
	fmt.Println("str=asdfghjkl result=", isUniqueChar("asdfghjkl"))
}

func isUniqueChar(s string) bool {
	arrRune := []rune(s)
	sort.Slice(arrRune, func(i, j int) bool {
		return arrRune[i] < arrRune[j]
	})

	for i := 0; i < len(arrRune)-1; i++ {
		if arrRune[i] == arrRune[i+1] {
			return false
		}
	}
	return true
}
