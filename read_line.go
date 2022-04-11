package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("c.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner()
}
