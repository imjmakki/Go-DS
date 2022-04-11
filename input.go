package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%T\n", scanner)

	scanner.Scan()

	text := scanner.Text()
}
