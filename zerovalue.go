package main

import "fmt"

func main() {

	var a = 4
	var b = 5.9

	a = int(b)
	fmt.Println(a, b)

	var value int
	var price float64
	var name string
	var done bool
	// zero values
	fmt.Println(value, price, name, done)
}
