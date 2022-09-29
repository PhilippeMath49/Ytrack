package main //Write a program that prints the name of the program.

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
}
