package main

import "fmt"

func f1() {
	fmt.Println("This is f1() function.")
}

func f2(a int, b int) {
	fmt.Println("Sum:", a+b)
}

func f3(a, b, c int, d, e float64, s string) {
	fmt.Println(a, b, c, d, e, s)
}

func f4(a float64) float64 {

}

func main() {
	f1()
	f2(5, 7)
	f3(4, 5, 6, 4.4, 7.8, "Golang")
}
