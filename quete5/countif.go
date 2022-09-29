package main

import (
	"fmt"
)

func main() {
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This", "1", "is", "4", "you"}
	answer1 := CountIf(IsNumeric, tab1)
	answer2 := CountIf(IsNumeric, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)
}

func IsNumeric(str string) bool {
	for _, v := range str {
		if v < '0' || v > '9' {

			return false
		}
	}
	return true
}

func CountIf(f func(string) bool, tab []string) int {
	result := make([]bool, len(tab))
	for i, v := range tab {
		result[i] = f(v)
	}
	count := 0
	for _, v := range result {
		if v == true {
			count++
		}
	}
	return count
}
