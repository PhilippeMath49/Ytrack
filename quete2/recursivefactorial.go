package main

import "fmt"

func main() {
	fmt.Println(recufa(6))

}

func recufa(number int) int {
	if number < 0 {
		return 0
	} else if number == 0 {
		return 1

	} else {
		return number * recufa(number-1)

	}

}
