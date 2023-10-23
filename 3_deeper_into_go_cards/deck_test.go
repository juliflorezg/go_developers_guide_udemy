//we need to initialize a go module, for this in this package directory run the command :
// go mod init cards
// that will create a go.mod file for us

package main

import "testing"

// when thinking about test, think 'what ares some easy assertions' or 'what do I care about with the code that I'm writing'

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// test that the deck actually contains 16 cards (according to the function)
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got: %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades', but got: %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be 'Four of Clubs', but got: %v", d[len(d)-1])
	}
}
