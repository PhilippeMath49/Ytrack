package main //Write a function that receives a string and a separator and returns a slice of strings that results of splitting the string s by the separator sep.

import (
	"fmt"
)

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}

func Split(s, sep string) []string {
	tab := []string{}
	word := ""*
	b:= 
	for _, v := range s {
		

		