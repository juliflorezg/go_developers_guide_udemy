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

	// jimPointer := &jim //> the & operator gives us the memory address of the value this variable is pointing at
	// fmt.Println(jimPointer)
	jim.print()
	//? in go we can call a pointer receiver with just a variable of that type, it is not necessary to first create a pointer and then call the function from that pointer, go will do that for us
	jim.updateName("Jimmy")
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

//> in Go this types are pass by value, so we have to use pointers to change the actual value in a function: int, float, string, bool, structs

//> these other types are pass by reference so we don't have to worry with pointers to change these values: slices, maps, channels, pointers, functions
