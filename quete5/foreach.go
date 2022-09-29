package main

import "fmt"

func main() {
	ForEach([]int{1, 2, 3, 4, 5}, func(n int) {
		fmt.Print(n)
	})
}

func ForEach(slice []int, f func(int)) {
	for _, v := range slice {
		f(v)
	}
}
