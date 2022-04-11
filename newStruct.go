package main

import "fmt"

func main() {
	type book struct {
		title, author string
		year          int
	}

	lastBook := book{title: "Anna karenina"}
	fmt.Println(lastBook.title)
	fmt.Printf("%#v\n", lastBook)

	lastBook.author = "leo Tolstoy"
	lastBook.year = 1878
	fmt.Printf("%+v\n", lastBook)

	aBook := book{title: "Anna karenina", author: "", year: 0}
	fmt.Println(aBook == lastBook)
}
