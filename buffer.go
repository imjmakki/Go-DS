package main

import "fmt"

func main() {
	c := make(chan int, 3)
	go func(c chan int) {
		fmt.Println("")
		c <- 10
		fmt.Println("")
	}(c)
}
