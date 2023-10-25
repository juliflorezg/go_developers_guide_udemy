package main

import "fmt"

// any type that has a function getGreeting that returns a string, now is going to be also of type bot, so we can use just one function for all those types (line 20)
type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// if we don't use the receiver value inside the function, we can delete it from the func declaration, like this: func (englishBot) getGreeting()...
func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// VERY custom logic for generating an spanish greeting
	return "Hola!"
}
