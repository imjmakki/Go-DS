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
}
