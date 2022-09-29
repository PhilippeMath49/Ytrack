package main

import "fmt"

func main() {
	fmt.Println(IsPrintable("hello"))
	fmt.Println(IsPrintable("hello\n"))
}

func IsPrintable(s string) bool {
	runeArray := []rune(s)
	for _, i := range runeArray {
		if charIsPrintable(i) {
			return false
		}
	}
	return true
}

func charIsPrintable(r rune) bool {
	if r >= 0 && r <= 31 {
		return true
	}
	return false
}
