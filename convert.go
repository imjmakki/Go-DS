package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x = 3
	var y = 3.1

	x = x * int(y)

	fmt.Println(x)

	s := string(99)
	fmt.Println(s)

	var myStr = fmt.Sprintf("%f", 44.2)
	fmt.Println(myStr)

	s1 := "3.123"

	var f1, err = strconv.ParseFloat(s1, 64)
	_ = err
	fmt.Println(f1)
}
