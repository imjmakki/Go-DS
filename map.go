package main

import "fmt"

func main() {
	f := fmt.Printf
	var employees map[string]string
	f("%#v\n", employees)
	f("No. of pairs : %d\n", len(employees))
	f("The value for key %q is %q\n", "Dan", employees["Dan"])

	var accounts map[string]float64
	f("%#v\n", accounts["0x323"])

	var m1 map[[5]int]string
	_ = m1

	people := map[string]float64{}

	people["John"] = 21.4
	people["Mary"] = 10
}
