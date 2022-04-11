package main

import "fmt"

func main() {
	friends := map[string]int{"Dan": 40, "Maria": 25}
	neighbors := friends
	fmt.Println(neighbors)

	friends["Dan"] = 50
	fmt.Println(neighbors)
}
