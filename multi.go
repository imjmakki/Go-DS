package main

import "fmt"

func main() {
	car, cost := "Audi", 50000
	fmt.Println(car, cost)

	car, year := "BMW", 2022
	_ = year

	opened, file := true, "a.txt"
	_, _ = opened, file
}
