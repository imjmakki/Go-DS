package main

import "fmt"

func main() {
	var cities []string
	fmt.Println("The cities is equal to nil", cities == nil)
	fmt.Printf("cities %#v\n", cities)
	fmt.Println(len(cities))
}
