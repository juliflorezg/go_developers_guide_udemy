package main

import "fmt"

func main() {
	// var card string = "Ace of spades"
	card := newCard()
	// card = "Five of spades "

	fmt.Println(card)
}

func newCard() string {
	return "Five of diamonds"
}
