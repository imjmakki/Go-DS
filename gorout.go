package main

import (
	"fmt"
	"runtime"
)

func l1() {
	fmt.Println("l1 (goroutine) execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("l1, i=", i)
	}
	fmt.Println("l1 execution finished")
}

func l2() {
	fmt.Println("l2 (goroutine) execution started")
	for i := 5; i < 16; i++ {
		fmt.Println("l2, i=", i)
	}
	fmt.Println("l2 execution finished")
}

func main() {
	fmt.Println("main execution started")
	fmt.Println("No. of CPU's:", runtime.NumCPU())
	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

	go l1()
	fmt.Println("No. of Goroutines after go l1():", runtime.NumGoroutine())

	l2()
	fmt.Println("main execution stopped")
}
