package main

import "fmt"

func main() {
	var1, var2 := 'a', 'b'
	fmt.Printf("Type: %T, Value: %d\n", var1, var2)

	str := "tara"
	fmt.Println(len(str))

	fmt.Println("Byte (not rune) at position 1:", str[1])
}
