package main

import "fmt"

func main() {
	//map type
	balances := map[string]float64{
		"USD": 2332.2,
		"EUR": 511.11,
	}
	fmt.Printf("%T\n", balances)

	//struct type
	type Person struct {
		name string
		age  int
	}

	var you Person
	fmt.Printf("%T\n", you)
}
