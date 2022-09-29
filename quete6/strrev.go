package main

import (
	"fmt"
)

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}

func StrRev(s string) string {
	var str string
	for _, c := range s {
		str = string(c) + str
	}
	return str
}
