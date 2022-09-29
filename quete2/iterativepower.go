package main

import "fmt"

func main() {
	fmt.Println(IterativePower(3, 3))
}

func IterativePower(nb int, power int) int {
	if power == 0 {
		return 1
	} else if nb == 0 {
		return 0
	} else {
		k := nb
		for i := 1; i < power; i++ {
			k = k * nb
		}
		return k
	}

}
