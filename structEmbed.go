package main

import "fmt"

func main() {
	type Contact struct {
		email, address string
		phone          int
	}
	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	john := Employee{
		name:   "John Keller",
		salary: 4000,
		contactInfo: Contact{
			email:   "jkeller@company.app",
			address: "Street 20, London",
			phone:   003233742223,
		},
	}
	fmt.Printf("%+v\n", john)
}
