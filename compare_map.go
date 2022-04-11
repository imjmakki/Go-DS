package main

import "fmt"

func main() {
	a := map[string]string{"X": "Y"}
	b := map[string]string{"A": "Y"}

	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)

	fmt.Println(s1, s2)
}
