package main

import (
	"fmt"
	"runtime"
)

func l1() {
	fmt.Println("l1 (goroutine) execution started")

	fmt.Println("l1 execution finished")
}

func main() {
	fmt.Println("main execution started")
	fmt.Println("No. of CPU's:", runtime.NumCPU())
	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
}
