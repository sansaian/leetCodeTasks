package main

import "fmt"

func main() {
	fmt.Println("string=`A man, a plan, a canal; Panama`  result", isPalindrome("A man, a plan, a canal; Panama"))
	fmt.Println("string=`A man, a plan, a canal: Panama`  result", isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println("string=`ZZBBAAAABbZz`  result", isPalindrome("ZZBBAAAABbZz"))
	fmt.Println("string=`race a car`  result", isPalindrome("race a car"))
	fmt.Println("string=` `  result", isPalindrome(" "))
	fmt.Println("string=`0P`  result", isPalindrome("0P"))
	fmt.Println("string=`a.`  result", isPalindrome("a."))
	fmt.Println("string=`\"7Ci`rd,9X;;r9,dX`iC7\" `  result", isPalindrome("\"7Ci`rd,9X;;r9,dX`iC7\""))
	fmt.Println("string=`\"7Ci`rd,9X;;r9,dX`iC7\" `  result", isPalindrome("7Ci`rd,9X;;r9,dX`iC7"))

}

func isPalindrome(s string) bool {
	runeSym := []rune(s)
	var j, k = 0, len(s) - 1
	for j < k {
		if runeSym[j] > 96 && runeSym[j] < 123 {
			runeSym[j] -= 32
			continue
		}
		if runeSym[k] > 96 && runeSym[k] < 123 {
			runeSym[k] -= 32
			continue
		}
		switch {
		case (runeSym[j] < 48 || runeSym[j] > 90) || (runeSym[j] > 57 && runeSym[j] < 65):
			j++
		case (runeSym[k] < 48 || runeSym[k] > 90) || (runeSym[k] > 57 && runeSym[k] < 65):
			k--
		case runeSym[j] == runeSym[k]:
			j++
			k--
		case runeSym[j] != runeSym[k]:
			return false
		}
	}
	return true
}
