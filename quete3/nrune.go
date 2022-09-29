package main

import (
	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Bye!", -2))
	z01.PrintRune(NRune("Bye!", 2))
	z01.PrintRune(NRune("Ola!", 2))
	z01.PrintRune('\n')
}

func NRune(s string, n int) rune {
	for i:=[]rune(s){
	if n < 0 {

		return 0
	} else if n > len(s) {
		return 0
	} else {
		i := []rune(s)
		return i[n-1]
	}
}
}