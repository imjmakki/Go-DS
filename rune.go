package main

import "fmt"

func main() {
	var1, var2 := 'a', 'b'
	fmt.Printf("Type: %T, Value: %d\n", var1, var2)

	str := "tara"
	fmt.Println(len(str))

	fmt.Println("Byte (not rune) at position 1:", str[1])

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}
}
