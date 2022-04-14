package main

import "fmt"

type car struct {
	brand string
	price int
}

func changeCar(c car, newBrand string, newPrice int) {
	c.price = newPrice
	c.brand = newBrand
}

func (c car) changeCar1(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

func main() {
	myCar := car{brand: "Audi", price: 40000}
	changeCar(myCar, "Renault", 20000)
	fmt.Println(myCar)

	myCar.changeCar1("Opel", 21000)
}
