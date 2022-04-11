package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("myFile.txt", os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bufferedWriter := bufio.NewWriter(file)
	bs := []byte{97, 98, 99}
}
