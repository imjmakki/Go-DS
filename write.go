package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("c.txt", os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bufferedWriter := bufio.NewWriter(file)
	bs := []byte{97, 98, 99}

	bytesWritten, err := bufferedWriter.Write(bs)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Byres written to buffer (not file) %d\n", bytesWritten)

	bytesAvailable := bufferedWriter.Available()
	log.Printf("Bytes available in buffer: %d\n", bytesAvailable)

	bytesWritten, err = bufferedWriter.WriteString("\nJust a random string")
	if err != nil {
		log.Fatal(err)
	}

	unFlushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unFlushedBufferSize)

	bufferedWriter.Flush()
}
