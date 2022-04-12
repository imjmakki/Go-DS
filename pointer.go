package main

import "fmt"

func main() {
	name := "Andrei"
	fmt.Println(&name)

	var x int = 2
	ptr := &x
	fmt.Printf("ptr is of type %T with a value of %v\n", ptr, ptr)
	fmt.Printf("address of x is %p\n", &x)
}
