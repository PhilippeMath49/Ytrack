package main

import (
	"fmt"
)

func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}

func Index(s string, toFind string) int {
	toFindCount := boucle(toFind)
	secondIndex := 0
	for _, j := range toFind {
		for i1, i2 := range s {
			if i2 == j {
				if toFindCount == 1 {
					return i1
				} else if toFindCount > 1 {
					for a := 0; a < toFindCount; a++ {
						if s[i1+a] == toFind[a] {
							secondIndex++
						} else {
							return -1
						}
					}
					if secondIndex == toFindCount {
						return i1
					}
				} else {
					return -1
				}
			}
		}
		if secondIndex <= 0 {
			return -1
		}
	}
	return toFindCount
}

func boucle(s string) int {
	count := 0
	for i := range s {
		count = i + 1
	}
	return count
}
