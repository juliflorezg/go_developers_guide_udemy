package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	var alex person

	fmt.Println(alex)
	fmt.Printf("%+v\n", alex) // %+v will print out all the different field names and values -> {firstName: lastName:}

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Printf("%+v\n", alex)
}
