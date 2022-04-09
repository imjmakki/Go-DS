package main

import "fmt"

func main() {
	var numbers [4]int

	fmt.Printf("%v\n", numbers)
	fmt.Printf("%#v\n", numbers)

	var a1 = [4]float64{}
	fmt.Printf("%#v\n", a1)
}
