package main

import "fmt"

func main() {
	c := make(chan int)
	c1 := make(<-chan string)
	c2 := make(chan<- string)

	fmt.Println("%T", "%T", "%T", c, c1, c2)
}
