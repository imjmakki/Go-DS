package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"b.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE
			0644,
	)
	if err != nil{
		log.Fatal(err)
	}
}
