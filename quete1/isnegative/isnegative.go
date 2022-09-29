package main

import "github.com/01-edu/z01"

func main() {
	Isnegative(0)
}
func Isnegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
}
