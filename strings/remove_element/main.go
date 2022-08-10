package main

import "fmt"

func main() {
	fmt.Println(removeDigit("ывфвы", 'в'))
	fmt.Println(removeDigit("133235", '3'))
}

func removeDigit(str string, digit rune) string {
	numberByte := []rune(str)
	for i := 0; i < len(numberByte); i++ {
		if numberByte[i] == digit {
			numberByte = append(numberByte[:i], numberByte[i+1:]...)
			return string(numberByte)
		}
	}
	return str
}
