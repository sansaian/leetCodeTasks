package main

import "fmt"

func main() {
	fmt.Println(maxVowels("abciiidef", 3)) //3
	fmt.Println(maxVowels("aeiou", 2))     //2
	fmt.Println(maxVowels("leetcode", 3))  //2
	fmt.Println(maxVowels("weallloveyou", 7))
	fmt.Println(maxVowels("rhythms", 4))                             //0
	fmt.Println(maxVowels("tryhard", 4))                             //1
	fmt.Println(maxVowels("ibpbhixfiouhdljnjfflpapptrxgcomvnb", 33)) //7
	fmt.Println(maxVowels("ramadan", 2))                             //1

}

func maxVowels(s string, k int) int {
	var r, l, sumVow, result int
	for ; r < len(s); r++ {
		if isVowel(s[r]) && sumVow < k {
			sumVow++
		}
		for r-l >= k {
			if isVowel(s[l]) {
				sumVow--
			}
			l++
		}
		result = max(result, sumVow)
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func isVowel(ch byte) bool {
	switch ch {
	case 'a':
		return true
	case 'e':
		return true
	case 'i':
		return true
	case 'o':
		return true
	case 'u':
		return true
	}
	return false
}
