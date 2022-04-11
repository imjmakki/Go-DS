package main

import "fmt"

func main() {
	type book struct {
		title, author string
		year          int
	}

	lastBook := book{title: "Anna karenina"}
	fmt.Println(lastBook.title)
}
