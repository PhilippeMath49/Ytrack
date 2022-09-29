package main

import "fmt"

func main() {
	a1 := []string{"Hello", "how", "are", "you"}
	a2 := []string{"This", "is", "4", "you"}

	result1 := Any(IsNumeric, a1)
	result2 := Any(IsNumeric, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}

func IsNumeric(str string) bool {
	for _, v := range str {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}

func Any(f func(string) bool, a []string) bool {
	result := make([]bool, len(a))
	for i, v := range a {
		result[i] = f(v)
	}
	for _, v := range result {
		if v == true {
			return true
		}
	}
	return false
}
