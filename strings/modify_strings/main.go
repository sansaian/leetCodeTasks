package main

import "fmt"

func main() {
	// think about ascii or unicode? если один символ больше чем 1 байт это работать не будет
	fmt.Println(isOneModify("pale", "bakebs"))
	fmt.Println(isOneModify("pale", "ple"))
	fmt.Println(isOneModify("pales", "pale"))
	fmt.Println(isOneModify("pale", "bale"))
	fmt.Println(isOneModify("pale", "bake"))
	fmt.Println(isOneModify("paable", "paale"))
}

func isOneModify(first, second string) bool {
	if len(first) == len(second) {
		return oneEditReplace(first, second)
	} else if len(first)+1 == len(second) {
		return oneEditInsert(first, second)
	} else if len(first)-1 == len(second) {
		return oneEditInsert(second, first)
	}
	return false
}

func oneEditInsert(first string, second string) bool {
	var index1, index2 int

	for index2 < len(second) && index1 < len(first) {
		if first[index1] != second[index2] {
			if index1 != index2 {
				return false
			}
			index2++
		} else {
			index1++
			index2++
		}
	}
	return true
}

func oneEditReplace(first, second string) bool {
	foundDiff := false
	for i := range first {
		if first[i] != second[i] {
			if foundDiff {
				return false
			}
			foundDiff = true
		}
	}
	return true
}
