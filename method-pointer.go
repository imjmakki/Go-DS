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

func (c *car) changeCar2(newBrand string, newPrice int) {
	(*c).brand = newBrand
	(*c).price = newPrice
}

func main() {
	myCar := car{brand: "Audi", price: 80000}
	changeCar(myCar, "Renault", 25000)
	fmt.Println(myCar)

	myCar.changeCar1("Opel", 15000)
	fmt.Println(myCar)

	myCar.changeCar2("Seat", 10000)
	fmt.Println(myCar)

	(&myCar).changeCar2("Ferrari", 250000)
	fmt.Println(myCar)

	var yourCar *car
	yourCar = &myCar
	fmt.Println(*yourCar)

	yourCar.changeCar2("VW", 22000)
	fmt.Println(*yourCar)

	(*yourCar).changeCar2("Bugatti", 2000000)
	fmt.Println(*yourCar)

	fmt.Println(myCar)
}
