package main

import "fmt"

type emptyInterface interface {
}

type person struct {
	info interface{}
}

func main() {
	var empty interface{}

	empty = 5
	fmt.Println(empty)

	empty = "Golang"
	fmt.Println(empty)

	empty = []int{5, 5, 6}
	fmt.Println(empty)
	fmt.Println(len(empty.([]int)))

	you := person{}
}
