package main

import "fmt"

func g1(a ...int) {
	fmt.Printf("%T\n", a)
	fmt.Printf("%#v\n", a)
}

func main() {
	g1(1, 2, 3, 4)
	g1()

	nums := []int{1, 2}
	nums = append(nums, 3, 4)

	g1(nums...)
}
