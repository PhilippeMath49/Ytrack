package main

import "fmt"

func main() {
	fmt.Print(RecursivePower(2, 8))

}

func RecursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	} else if nb == 0 {
		return 0
	} else {
		return nb * RecursivePower(nb, power-1)
	}

}
