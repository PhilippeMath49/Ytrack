package main

import (
	"fmt"
)

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}

func ConcatParams(args []string) string {
	str := ""
	for _, v := range args {
		str += v + "\n"

	}
	return str
}
