package main

import "fmt"

func g1(a ...int) {
	fmt.Printf("%T\n", a)
	fmt.Printf("%#v\n", a)
}

func main() {

	g1()
}
