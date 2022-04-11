package main

import "fmt"

func main() {
	f := fmt.Printf
	var employees map[string]string
	f("%#v\n", employees)
	f("No. of pairs : %d\n", len(employees))
	f("The value for key %q is %q\n", "Dan", employees["Dan"])

}
