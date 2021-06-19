package main

import "fmt"

func main() {
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()           // startIndex to 4
	// remainingCards.print() // 5 to lastIndex
	// cards.saveToFile("my_cards")
	newDeck := newDeckFromFile("my_cards")
	newDeck.print()
	fmt.Println("-----Shuffle-----")
	newDeck.shuffle()
	newDeck.print()
}
