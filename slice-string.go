package main

import "fmt"

func main() {
	s1 := "I love golang!"
	fmt.Println(s1[2:5])

	rs := []rune(s1)
	fmt.Printf("%T\n", rs)
}
