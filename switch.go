package main

import (
	"fmt"
	"time"
)

func main() {
	language := "Golang"

	switch language {
	case "Python":
		fmt.Println("you are learning Python! You don't use curly braces but indentation!!")
		break
	case "Go", "Golang":
		fmt.Println("Great, go for Go :) You are using curley braces{} :)")
		break
	default:
		fmt.Println("C/C++ is the basic of programming :)")
		break
	}

	n := 5
	switch true {
	case n%2 == 0:
		fmt.Println("Even number :)")
	case n%2 != 0:
		fmt.Println("Odd number (:")
	default:
		fmt.Println("Never been here")
	}

	hour := time.Now().Hour()
	fmt.Println(hour)

	switch {
	case hour < 12:
		fmt.Println("Good Morning :)")
	case hour < 17:
		fmt.Println("Good Afternoon :)")
	default:
		fmt.Println("Good Evening :)")
	}
}
