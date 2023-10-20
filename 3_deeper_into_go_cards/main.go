package main

// go run main.go deck.go
// to run the program with all the package's files
func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

}
