package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := Map(IsPrime, a)
	fmt.Println(result)
}

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))
	for i, v := range a {
		result[i] = f(v)
	}
	return result

}

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	for x := 2; x < nb; x++ {
		if nb%x == 0 {
			return false
		}

	}
	return true
}
