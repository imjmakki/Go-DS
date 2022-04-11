package main

import "fmt"

func main() {
	f := fmt.Printf
	var employees map[string]string
	f("%#v\n", employees)

}
