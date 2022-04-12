package main

import "fmt"

func main() {
	a := 5.5
	p1 := &a
	pp1 := p1
	fmt.Printf("Value of p1: %v, address of p1: %v\n", p1, &p1)
}
