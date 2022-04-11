package main

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
}
