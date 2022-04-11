package main

import "fmt"

func f1() {
	fmt.Println("This is f1() function.")
}

func f2(a int, b int) {
	fmt.Println("Sum:", a+b)
}

func main() {
	f1()
	f2(5, 7)
}
