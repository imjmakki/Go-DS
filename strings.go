package main

import "fmt"

func main() {
	s1 := "Hi there Go!"
	fmt.Println(s1)

	fmt.Println("He Says: \"Hello\"")
	fmt.Println(`He Says: "Hello"`)

	s2 := `I like \n Go!`
	fmt.Println(s2)

	fmt.Println("Price: 100\nBrand: Nike")
	fmt.Println(`
Price: 100
Brand: Nike
`)

	fmt.Println(`C:\Users\MJ`)
	fmt.Println("C:\\Users\\MJ")
}
