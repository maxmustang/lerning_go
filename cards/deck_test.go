package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("expected deck length of 20, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("expeceted first card Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("expected last card to be Four of Clubs, but got %v", d[len(d)-1])
	}

}

func xpo(t *testing.T) {

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	const _decktesting string = "_decktesting"
	os.Remove(_decktesting)
	d := newDeck()
	d.saveToFile(_decktesting)

	loadedDeck := newDeckFromFile(_decktesting)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove(_decktesting)

}
