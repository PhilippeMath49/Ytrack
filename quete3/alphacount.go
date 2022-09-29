package main

import (
	"fmt"
)

func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}

func AlphaCount(s string) int {
	count := 0
	for _, i := range s {
		if isAlpha(i) {
			count++
		}

	}
	return count
}

func isAlpha(k rune) bool {
	for a := 'a'; a <= 'z'; a++ {
		if k == a {
			return true
		}
	}
	for a := 'A'; a <= 'Z'; a++ {
		if k == a {
			return true
		}
	}
	return false
}
