// package is a project or workspace
// we name the package 'main' in order to get an executable package
package main

import "fmt"

// use go run main.go to run this program
// go build allow us to compile the program to machine code
// go fmt formats all the code in each file in the current directory
// go install compiles and install dependencies
// go get download raw source code of someone else's package
// go test any test associated with this project

func main() {
	fmt.Println("Hello World!")
}
