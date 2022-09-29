package main

import (
	"fmt"
)

func main() {
	a := 13
	b := 2
	c := a
	UltimateDivMod(&a, &b, &c)
	fmt.Println(a)
	fmt.Println(b)
}

func UltimateDivMod(a *int, b *int, c *int) {

	*a = *a / *b
	*b = *c % *b
}
