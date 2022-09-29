package main

import (
	"fmt"
)

func main() {
	fmt.Println(BasicAtoi("123045"))
	fmt.Println(BasicAtoi("00008000012345"))
	fmt.Println(BasicAtoi("000000"))
}

func BasicAtoi(s string) int {
	var result int
	for _, char := range s {
		result = result*10 + int(char-'0')
	}
	return result
}
