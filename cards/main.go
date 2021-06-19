package main

func main() {
	cards := deck{"a card", newCard()}
	cards = append(cards, "another card") // return a new slice (not modify existing ones)
	cards.print()
}

func newCard() string {
	return "Ace of Spade"
}
