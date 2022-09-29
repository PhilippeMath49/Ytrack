package main

import "fmt"

func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}

func SplitWhiteSpaces(str string) []string {
	tab := []string{}
	word := ""
	for _, v := range str {
		if v == ' ' {
			tab = append(tab, word)
			word = ""

		} else {
			word += string(v)
		}
	}
	tab = append(tab, word)
	return tab

}
