package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int, 3)
	_ = c2
	go func(c chan int) {
		fmt.Println("func goroutine starts sending data into the channel")
		c <- 10
		fmt.Println("func goroutine after sending data into the channel")
	}(c1)
	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)
}
