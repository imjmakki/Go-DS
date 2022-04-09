package main

import "fmt"

func main() {
	var numbers [4]int

	fmt.Printf("%v\n", numbers)
	fmt.Printf("%#v\n", numbers)

	var a1 = [4]float64{}
	fmt.Printf("%#v\n", a1)

	var a2 = [4]int{-10, 1, 100}
	fmt.Printf("%#v\n", a2)

	a3 := [4]string{"Dan", "Diana", "Paul", "John"}
	fmt.Printf("%#v\n", a3)
}
