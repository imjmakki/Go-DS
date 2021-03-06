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

	p(strings.ToUpper("go") == strings.ToUpper("Go"))
	p(strings.EqualFold("GO", "go"))

	myStr := strings.Replace("192.168.0.1", ".", ":", 2)
	p(myStr)

	myStr = strings.Replace("192.168.0.1", ".", ":", -1)
	p(myStr)

	myStr = strings.ReplaceAll("192.168.0.1", ".", ":")
	p(myStr)

	s := strings.Split("a,b,c", ",")
	fmt.Printf("%T\n", s)
	fmt.Printf("%#v\n", s)

	s = strings.Split("Go for Go!", "")
	fmt.Printf("%#v\n", s)

	s = []string{"I", "Learn", "Golang"}
	myStr = strings.Join(s, " ")
	p(myStr)

	myStr = "Orange Green \n Blue Yellow"
	fields := strings.Fields(myStr)
	fmt.Printf("%T\n", fields)
	fmt.Printf("%#v\n", fields)

	s1 := strings.TrimSpace("\t  Goodbye Windows, Welcome Linux!  \n")
	fmt.Printf("%q\n", s1)

	s2 := strings.Trim("...Hello Gophers!!!!?", ".!?")
	fmt.Printf("%q\n", s2)
}
