package main

import (
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("b.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := make([]byte, 2)

	numberBytesRead, err := io.ReadFull(file, byteSlice)
}
