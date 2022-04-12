package main

import "fmt"

func main() {
	a := 5.5
	p1 := &a
	pp1 := p1
	fmt.Printf("Value of p1: %v, address of p1: %v\n", p1, &p1)
	fmt.Printf("Value of pp1: %v, address of pp1: %v\n", pp1, &pp1)

	fmt.Printf("*p is %v\n", *p1)
	fmt.Printf("*pp1 is %v\n", *pp1)
}
