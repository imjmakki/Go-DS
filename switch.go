package main

import "fmt"

func main() {
	language := "Golang"

	switch language {
	case "Python":
		fmt.Println("you are learning Python! You don't use curly braces but indentation!!")
		break
	}
}
