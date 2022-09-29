package main

import "fmt"

func main() {
	fmt.Println(FindNextPrime(9536525))

}

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	for x := 2; x < nb; x++ {
		if nb%x == 0 {
			return false
		}

	}
	return true
}

func FindNextPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	} else {
		n := nb + 1
		for IsPrime(n) == false {
			n++
		}
		return n
	}
}
