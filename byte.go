package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "Codruta"
	fmt.Println(len(name))

	n := utf8.RuneCountInString(name)
	fmt.Println(n)
}
