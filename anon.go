package main

import "fmt"

func main() {
	func(msg string) {
		fmt.Println(msg)
	}("I'm an anonymous function!:)")
}
