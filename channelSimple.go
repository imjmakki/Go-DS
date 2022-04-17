package main

import "fmt"

func x1(n int, ch chan int) {
	ch <- n
}

func main() {
	c := make(chan int)
	c1 := make(<-chan string)
	c2 := make(chan<- string)

	fmt.Println("%T", "%T", "%T", c, c1, c2)
	go x1(10, c)
	n := <-c
	fmt.Println("Value received:", n)
	fmt.Println("Exiting main ...")
}
