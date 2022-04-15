package main

import "math"

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func printCircle(c circle) {

}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {

}
