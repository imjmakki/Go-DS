package main

import "fmt"

func main() {
	//typed constant
	const days int = 7
	const (
		//untyped constant
		min1 = 500
		min2
		min3
	)
	fmt.Println(min1, min2, min3)

	const d = 5 > 10
	fmt.Println(d)

	const y = 2.2 * 5
	fmt.Printf("%T\n", y)

	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)
	fmt.Println(c1, c2, c3)

	const (
		c11 = iota
		c22
		c33
	)
	fmt.Println(c11, c22, c33)

	const (
		a = iota * 2
		b
		c
	)
	fmt.Println(a, b, c)
}
