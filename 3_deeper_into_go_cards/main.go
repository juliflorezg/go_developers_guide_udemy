package main

// go run main.go deck.go
// to run the program with all the package's files
func main() {
	cards := newDeck()

	cards.print()
	// fmt.Println(cards)
}
