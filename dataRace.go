package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const gr = 100
	var wg sync.WaitGroup
	wg.Add(gr * 2)

	var n int = 0
	var m sync.Mutex

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 2)
			m.Lock()
			n++
			m.Unlock()
			wg.Done()
		}()
		go func() {
			time.Sleep(time.Second / 2)
			m.Lock()
			defer m.Unlock()
			n--
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final value of n:", n)
}
