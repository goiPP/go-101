package main

import "fmt"

func main() {
	cards := []string{"a card", newCard()}
	cards = append(cards, "another card") // return a new slice (not modify existing ones)
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Ace of Spade"
}
