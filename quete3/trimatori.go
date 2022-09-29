package main

import "fmt"

func main() {

	fmt.Println(TrimAtoi("Hello World"))
}

func TrimAtoi(s string) int {

	arr := []rune(s)

	c := 0
	num_our := 0
	sign := 1
	minus := false

	for _, word := range arr {

		if word >= '0' && word <= '9' {
			for i := '0'; i < word; i++ {

				c++
			}
			num_our = num_our*10 + c
			c = 0

		}

		if word == '-' && num_our == 0 {
			minus = true
		}
	}
	if minus == true {
		sign = -1
	}

	return num_our * sign

}
