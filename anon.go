package main

import "fmt"

func increment(x int) func() int {

}

func main() {
	func(msg string) {
		fmt.Println(msg)
	}("I'm an anonymous function!:)")
}
