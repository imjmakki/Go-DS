package main

import "fmt"

func g1(a ...int) {
	fmt.Printf("%T\n", a)
	fmt.Printf("%#v\n", a)
}

func g2(a ...int) {
	a[0] = 50
}

func SumAndProduct(a ...float64) (float64, float64) {
	sum := 0
	product := 1
	for _, v := range a {
		sum += v
		product *= v
	}
}

func main() {
	g1(1, 2, 3, 4)
	g1()

	nums := []int{1, 2}
	nums = append(nums, 3, 4)

	g1(nums...)
	g2(nums...)
	fmt.Println("Nums:", nums)
}
