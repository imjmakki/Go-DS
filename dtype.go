package main

import "fmt"

func main() {
	//map type
	balances := map[string]float64{
		"USD": 232.2,
		"EUR": 511.11,
	}
	fmt.Printf("%T\n", balances)
}
