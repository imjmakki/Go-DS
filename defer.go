package main

import "fmt"

func foo() {
	fmt.Println("This is foo()")
}

func bar() {
	fmt.Println("This is bar()")
}

func foobar() {
	fmt.Println("This is foobar()")
}

func main() {
	defer foo()
	bar()
	fmt.Println("Just a string after defering foo() and calling bar()")
	defer foobar()
}
