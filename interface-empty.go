package main

import "fmt"

type emptyInterface interface {
}

func main() {
	var empty interface{}

	empty = 5
	fmt.Println(empty)

	empty = "Golang"
	fmt.Println(empty)

	empty = []int{1, 2, 3}
	fmt.Println(empty)
}
