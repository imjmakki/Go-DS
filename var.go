package main

import "fmt"

func main() {
	var age int = 26
	fmt.Println("Age: ", age)

	var name = "MJ"
	//fmt.Println("Your name is: ", name)
	_ = name // to avoiding errors
}
