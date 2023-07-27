package main

func main() {
	// cards := newDeck()

	// hand, ramainingCards := deal(cards, 5)

	// hand.print()
	// ramainingCards.print()

	cards := newDeck()
	//cards.saveToFile("cards-text")
	//cards := newDeckFromFile("cards-text")
	cards.shuffle()
	cards.print()
}
