package main

import "fmt"

func f1(n int, ch chan int) {
	ch <- n
}

func main() {
	c := make(chan int)
	c1 := make(<-chan string)
	c2 := make(chan<- string)

	fmt.Println("%T", "%T", "%T", c, c1, c2)
	go f1(10, c)
}
