package main

import "fmt"

func main() {
	fmt.Println(ToLower("Hello! How are you?"))
}

func ToLower(s string) string {
	runeArray := []rune(s)
	for i := range runeArray {
		if runeArray[i] >= 'A' && runeArray[i] <= 'Z' {
			runeArray[i] = runeArray[i] + 32
		}
	}
	return string(runeArray)
}
