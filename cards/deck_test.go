package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got  %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting") // to remove any dirty test files from previous test run

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got  %v", len(loadedDeck))
	}

	os.Remove("_decktesting") // to remove any dirty test files before end this test
}
