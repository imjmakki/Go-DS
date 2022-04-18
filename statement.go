package main

import "time"

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Hello"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Salute"
	}()
}
