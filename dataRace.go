package main

import (
	"sync"
	"time"
)

func main() {
	const gr = 100
	var wg sync.WaitGroup
	wg.Add(gr * 2)

	var n int = 0

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 2)
		}
	}
}
