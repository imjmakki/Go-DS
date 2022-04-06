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
}
