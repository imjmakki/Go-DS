package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("main execution started")
	fmt.Println("No. of CPU's:", runtime.NumCPU())
	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())
}
