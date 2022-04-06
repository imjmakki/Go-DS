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
	fmt.Println(f1 * 3.4)

	i, err := strconv.Atoi("-50")
	s2 := strconv.Itoa(20)

	fmt.Printf("i type is %T, i vlue is %v", i, i)
	fmt.Printf("s type is %T, i value is %q", s2, s2)
}
