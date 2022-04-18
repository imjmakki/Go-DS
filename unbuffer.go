package main

import "fmt"

func main() {
	c1 := make(chan int)
	c2 := make(chan int, 3)
	go func(c chan int) {
		fmt.Println("func goroutine starts sending data into the channel")
		fmt.Println("func goroutine after sending data into the channel")
	}()
}
