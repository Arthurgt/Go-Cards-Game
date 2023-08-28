package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 48 {
		t.Errorf("Expected deck length of 48, but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades' but got %v", d[0])
	}
	if d[len(d)-1] != "Jack of Clubs" {
		t.Errorf("Expected last card to be 'Jack of Clubs' but got %v", d[0])
	}
}

func TestSaveDeckToFileAndReadDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveDeckToFile("_decktesting")

	loadedDeck := readDeckFromFile("_decktesting")
	if len(loadedDeck) != 48 {
		t.Errorf("Expected 48 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
