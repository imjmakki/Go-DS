package main

import "fmt"

func changeValues(quantity int, price float64, name string, sold bool) {
	quantity = 3
	price = 500.4
	name = "Mobile Phone"
	sold = false
}

func main() {
	quantity, price, name, sold := 5, 300.4, "Laptop", true
	fmt.Println("", quantity, price, name, sold)
}
