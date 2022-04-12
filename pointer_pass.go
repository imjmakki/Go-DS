package main

import "fmt"

func change(a *int) *float64 {
	*a = 100

	b := 5.5
	return &b
}

func main() {
	x := 8
	k := &x

	fmt.Println("value of x before calling change():", x)
	change(k)
	fmt.Println("value of x after calling change():", k)
}
