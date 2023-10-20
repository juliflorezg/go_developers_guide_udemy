package main

// go run main.go deck.go
// to run the program with all the package's files
func main() {
	cards := deck{"Ace of diamonds", newCard()}
	cards = append(cards, "Six of spades")

	cards.print()
	// fmt.Println(cards)
}

func newCard() string {
	return "Five of diamonds"
}
