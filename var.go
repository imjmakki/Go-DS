package main

import "fmt"

func main() {
	var age int = 26
	fmt.Println("Age: ", age)

	var name = "MJ"
	//fmt.Println("Your name is: ", name)
	_ = name // to avoiding errors

	s := "Golang Course!"
	fmt.Println(s)

	name1 := "Nabeel"
	_ = name1

	var opened = false
	_ = opened
}
