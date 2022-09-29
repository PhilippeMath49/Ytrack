package main

import "fmt"

func main() {
	fmt.Println(sqrt(16))
}

func sqrt(nb int) int {
	for i := nb; i > 0; i-- {
		if nb == (i * i) {
			return i
		}
	}
	return 0
}
