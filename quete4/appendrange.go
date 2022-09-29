package main

import (
	"fmt"
)

func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}

func AppendRange(min, max int) []int {
	if min >= max {
		var tab []int
		return tab

	} else {
		tab := []int{}
		for i := min; i <= max; i++ {
			tab = append(tab, i)

		}
		return tab
	}

}
