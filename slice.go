package main

import "fmt"

func main() {
	var cities []string
	fmt.Println("The cities is equal to nil", cities == nil)
	fmt.Printf("cities %#v\n", cities)
	fmt.Println(len(cities))

	numbers := []int{2, 3, 4, 5}
	fmt.Println(numbers)

	nums := make([]int, 2)
	fmt.Printf("%#v\n", nums)

	type name []string
}
