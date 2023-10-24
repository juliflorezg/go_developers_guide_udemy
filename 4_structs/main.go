package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@example.com",
			zipCode: 45456,
		},
	}

	jimPointer := &jim //> the & operator gives us the memory address of the value this variable is pointing at
	fmt.Println(jimPointer)
	jim.print()
	jimPointer.updateName("Jimmy")
	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// *person would be read as 'a pointer that points at a person'
func (pointerToPerson *person) updateName(newFirstName string) {
	//> the * operator gives us the value this memory address is pointing at
	(*pointerToPerson).firstName = newFirstName
}
