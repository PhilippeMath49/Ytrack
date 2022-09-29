package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsLower("HELLO"))
	fmt.Println(IsLower("HELLO!"))

}

func IsLower(str string) bool {
	runeArray := []rune(str)
	runeCount := arrayCount(str)
	count := 0
	for _, i := range runeArray {
		if i >= 'a' && i <= 'z' {
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
