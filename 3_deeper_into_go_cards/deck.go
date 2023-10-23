package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck'
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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// toString function converts a deck of cards to a string
func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	// 0666 means anyone can read or write this file
	return os.WriteFile(filename, []byte(d.toString()), 0666)

}

func newDeckFromFile(filename string) deck {
	// bs is for byte slice
	bs, err := os.ReadFile(filename)

	if err != nil {
		// Option 1: log the error and return a call to newDeck()
		// Option 2: log the error and entirely quit the program

		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",") // Ace of spades,Two of spades,...

	return deck(s)
}

// shuffle function shuffles the deck
func (d deck) shuffle() {
	for i := range d {

		rand.Seed(time.Now().UnixNano()) // seeding a random value is deprecated as of Go 1.21.0
		newPosition := rand.Intn(len(d))

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
