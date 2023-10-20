package main

import "fmt"

//create a new type of 'deck'
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// receiver of the function -> any variable of type 'deck' now gets access to the "print" method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
