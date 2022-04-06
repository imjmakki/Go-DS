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
	fmt.Println("%T\n", y)
}
