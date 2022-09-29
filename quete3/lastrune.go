package main

import "github.com/01-edu/z01"

func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("image"))
	z01.PrintRune('\n')
}

func LastRune(s string) rune {
	i := []rune(s)
	return i[len(i)-1]
}
