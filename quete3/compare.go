package main

import (
	"fmt"
)

func main() {
	fmt.Println(Compare("Hello!", "Hello!"))
	fmt.Println(Compare("Salut!", "lut!"))
	fmt.Println(Compare("Ola!", "Ol"))
}

func Compare(a, b string) int {
	i := len(a) - len(b)
	if i == 0 {
		return 0
	} else if i < 0 {
		return -1
	} else {
		return 1
	}

}
