package main

import "fmt"

func main() {
	factor(6)
}

func factor(n int) {
	if n <= 0 {
		fmt.Println(0)

	}
	result := 1
	for i := 1; i <= n; i++ {
		result = result * i
		fmt.Println(result)

	}

}
