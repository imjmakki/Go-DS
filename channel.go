package main

import "fmt"

func main() {
	var ch chan int
	fmt.Println(ch)

	ch = make(chan int)
	fmt.Println(ch)

	c := make(chan int)
	c <- 10
}
