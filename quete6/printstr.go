package main

func main() {
	PrintStr("Hello World!")
}

func PrintStr(s string) {
	for _, c := range s {
		print(string(c))
	}
}
