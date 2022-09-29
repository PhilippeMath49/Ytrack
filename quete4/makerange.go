package main

import (
	"fmt"
)

func main() {
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}

func MakeRange(min, max int) []int {
	if min >= max {
		var tab []int
		return tab

	} else {
		tab := make([]int, max-min)
		for i := range tab {
			tab[i] = min + i
		}
		return tab
	}

}
