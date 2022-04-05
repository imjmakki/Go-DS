package main

import "fmt"

func main() {
	name := "MJ"
	fmt.Println("Hello, All I'm", name)

	a, b := 5, 7
	fmt.Println("Sum:", a+b, "Mean Value:", (a+b)/2)

	fmt.Printf("Your age is %d\n", 21)
	fmt.Printf("x is %d, y is %f\n", 5, 9)
	fmt.Printf("He says: \"Hello Go!\"\n")

	figure := "Circle"
	radius := 5
	pi := 3.14159

	_ = figure

	fmt.Printf("Radius is %d\n", radius)
	fmt.Printf("Radius is %+d\n", radius)

	fmt.Printf("Pi constant is %f\n", pi)
	fmt.Printf("The diameter of a %s with a Radius of %d is %f\n", figure, radius, float64(radius)*2*pi)

	fmt.Printf("This is a %q\n", figure)

	fmt.Printf("The diameter of a %v with a Radius of %v is %v\n", figure, radius, float64(radius)*2*pi)

	fmt.Printf("figure is of type %T \n", figure)
	fmt.Printf("radius is of type %T \n", radius)
	fmt.Printf("pi is of type %T \n", pi)
}
