package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("45")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
}
