package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("os.Args", os.Args)
	fmt.Println("Path: ", os.Args[0])
	fmt.Println("1st arguments: ", os.Args[1])
	fmt.Println("2nd arguments: ", os.Args[2])
	fmt.Println("No. of items inside os.Args: ", len(os.Args))

	var result, err = strconv.ParseFloat(os.Args[1], 64)
	fmt.Println(result)
	_ = err
}