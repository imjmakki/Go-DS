package main

import "sync"

func main() {
	const gr = 100
	var wg sync.WaitGroup
	wg.Add(gr * 2)
}
