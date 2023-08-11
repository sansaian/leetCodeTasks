package main

import "fmt"

func main() {
	fmt.Println(maxNumberOfBalloons("nlaebolko"))
}

func maxNumberOfBalloons(text string) int {
	store := make(map[byte]int, len(text))
	for i := 0; i < len(text); i++ {
		store[text[i]] = store[text[i]] + 1
	}
	result := 0
	template := "balloon"
	for {
		for i := 0; i < len(template); i++ {
			store[template[i]]--
			if store[template[i]] < 0 {
				return result
			}
		}
		result++
	}
}
