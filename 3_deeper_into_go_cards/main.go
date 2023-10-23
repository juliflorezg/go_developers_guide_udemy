package main

import "fmt"

// go run main.go deck.go
// to run the program with all the package's files
func main() {
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	cards := newDeck()
	// cards.toString()
	fmt.Println(cards.toString())

	cards.saveToFile("my_cards")
}
