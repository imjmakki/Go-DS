package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var newFile *os.File
	fmt.Printf("%T\n", newFile)

	var err error
	newFile, err = os.Create("a.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		log.Fatalln(err)
	}

	err = os.Truncate("a.txt", 0)

	if err != nil {
		log.Fatal(err)
	}
	newFile.Close()

	file, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	file1, err := os.OpenFile("a.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	file1.Close()

	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("a.txt")
	p := fmt.Println
}
