package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))

}
func IsAlpha(str string) bool {
	runeArray := []rune(str)
	runeCount := arrayCount(str)
	count := 0
	for _, char := range runeArray {
		if isNumberOrAlph(char) {
			count++
		}
	}
	if count == runeCount {
		return true
	}
	return false
}

func arrayCount(s string) int {
	count := 0
	for i := range s {
		count = i + 1
	}
	return count
}
func isNumberOrAlph(r rune) bool {
	if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
		return true
	}
	return false
}
