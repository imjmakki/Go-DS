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

	scanner := bufio.NewScanner(file)
	success := scanner.Scan()
	if success == false {
		err = scanner.Err()
	}
}
