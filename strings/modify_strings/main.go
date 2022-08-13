package main

import "fmt"

func main() {
	fmt.Println(isOneModify("pale", "bakebs"))
	fmt.Println(isOneModify("pale", "ple"))
	fmt.Println(isOneModify("pales", "pale"))
	fmt.Println(isOneModify("pale", "bale"))
	fmt.Println(isOneModify("pale", "bake"))
	fmt.Println(isOneModify("paable", "paale"))
}

func isOneModify(a, b string) bool {
	if len(a)-len(b) >= 2 || len(b)-len(a) >= 2 {
		return false
	}
	diff := 0
	setA := make(map[rune]int, len(a))
	setB := make(map[rune]int, len(b))
	for _, sym := range a {
		setA[sym]++
	}
	for _, sym := range b {
		setB[sym]++
	}
	for sym := range setA {
		if setA[sym] != setB[sym] {
			diff++
		}

	}
	if len(a) < len(b) {
		diff++
	}
	return diff <= 1
}
