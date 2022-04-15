package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func p1() {
	fmt.Println("p1 (goroutine) execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("p1, i=", i)
	}
	fmt.Println("p1 execution finished")
}

func p2() {
	fmt.Println("p2 (goroutine) execution started")
	for i := 5; i < 16; i++ {
		fmt.Println("p2, i=", i)
	}
	fmt.Println("p2 execution finished")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go p1()
	fmt.Println("No. of Goroutines after go p1():", runtime.NumGoroutine())

	p2()
	time.Sleep(time.Second * 2)
	fmt.Println("main execution stopped")
}
