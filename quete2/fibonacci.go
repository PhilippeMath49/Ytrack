package main

import "fmt"

func main() {
	fmt.Print(Fibonacci(2))

}

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}
	if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	}
	return Fibonacci(index-1) + Fibonacci(index-2)
}
