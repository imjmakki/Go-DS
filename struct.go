package main

import "fmt"

func main() {
	title1, author1, year1 := "The Devin Comedy", "Dante Alighieri", 1320
	title2, author2, year2 := "Macbeth", "William Shakespeare", 1606

	fmt.Println("Book1:", title1, author1, year1)
	fmt.Println("Book2:", title2, author2, year2)

	type book struct {
		title  string
		author string
		year   int
	}

	type book1 struct {
		title, author string
		year          int
	}

	myBook := book{"The Devin Comedy", "Dante Alighieri", 1320}
	fmt.Println(myBook)
}
