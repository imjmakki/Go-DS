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
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", numberBytesRead)
	log.Printf("Data read: %s\n", byteSlice)

	file, err = os.Open("read.go")
	if err != nil {
		log.Fatal(err)
	}
}
