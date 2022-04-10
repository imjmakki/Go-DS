package main

import (
	"fmt"
	"strings"
)

func main() {
	p := fmt.Println
	result := strings.Contains("I love Go Programming!", "love")
	p(result)

	result = strings.Contains("success", "xy")
	p(result)

	result = strings.ContainsRune("Golang", 'g')
	p(result)

	n := strings.Count("cheese", "e")
	p(n)

	n = strings.Count("Five", "")
	p(n)

	p(strings.ToLower("Go Python Java"))
	p(strings.ToUpper("Go Python Java"))
}
