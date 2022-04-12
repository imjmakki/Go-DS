package main

import "fmt"

func changeValues(quantity int, price float64, name string, sold bool) {
	quantity = 3
	price = 500.4
	name = "Mobile Phone"
	sold = false
}

func changeValuesByPointer(quantity *int, price *float64, name *string, sold *bool) {
	*quantity = 3
	*price = 500.4
	*name = "Mobile Phone"
	*sold = false
}

type Product struct {
	price       float64
	productName string
}

func changeProduct(p Product) {
	p.price = 300
	p.productName = "Bicycle"
}

func changeProductByPointer(p *Product) {
	(*p).price = 300
	(*p).productName = "Bicycle"
}

func changeSlice(s []int) {
	for i := range s {
		s[i]++
	}
}

func changeMap(m map[string]int) {
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
}

func main() {
	quantity, price, name, sold := 5, 300.4, "Laptop", true
	fmt.Println("Before calling changeValue():", quantity, price, name, sold)
	changeValues(quantity, price, name, sold)
	fmt.Println("After calling changeValue():", quantity, price, name, sold)

	changeValuesByPointer(&quantity, &price, &name, &sold)
	fmt.Println("After calling changeValueByPointer():", quantity, price, name, sold)

	gift := Product{
		price:       100,
		productName: "Watch",
	}
	changeProduct(gift)
	fmt.Println(gift)

	fmt.Println("Before calling the changeProductByPointer():", gift)
	changeProductByPointer(&gift)
	fmt.Println("After calling the changeProductByPointer():", gift)

	prices := []int{1, 2, 3}
}
