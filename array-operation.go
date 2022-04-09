package main

import "fmt"

func main() {
	numbers := [3]int{10, 20, 30}
	fmt.Printf("%#v\n", numbers)

	numbers[0] = 7
	fmt.Printf("%#v\n", numbers)

	for i, v := range numbers {
		fmt.Println("index:", i, "value:", v)
	}

	for i = 0; i <len(numbers), i++{

	}
}
