package main

import "fmt"

func main() {
	price, inStock := 100, true

	if price > 80 {
		fmt.Println("Too expensive!")
	}

	if price <= 100 && inStock == true {
		fmt.Println("Buy it!")
	}

	if price < 100 {
		fmt.Println("It's cheap!")
	} else if price == 100 {
		fmt.Println("On the edge!")
	} else {
		fmt.Println("It's expensive")
	}
}
