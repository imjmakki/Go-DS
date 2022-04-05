package main

import "fmt"

func main() {
	car, cost := "Audi", 50000
	fmt.Println(car, cost)

	car, year := "BMW", 2022
	_ = year

	opened, file := true, "a.txt"
	_, _ = opened, file

	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	var i, j int
	i, j = 5, 8
	j, i = i, j //swapping variables

	fmt.Println(i, j)
}
