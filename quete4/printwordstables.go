package main

import "github.com/01-edu/z01"

func main() {
	PrintWordsTables([]string{"Hello", "how", "are", "you?"})
}

func PrintWordsTables(table []string) {
	for _, v := range table {
		for _, r := range v {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
