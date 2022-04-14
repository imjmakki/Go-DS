package main

import (
	"fmt"
	"time"
)

type names []string

func (n names) pring() {

}

func main() {
	const day = 24 * time.Hour
	fmt.Printf("%T\n", day)

	seconds := day.Seconds()
	fmt.Printf("%T\n", seconds)
	fmt.Printf("Seconds in a day: %v\n", seconds)
}
