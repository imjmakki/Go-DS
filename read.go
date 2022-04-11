package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("b.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}
