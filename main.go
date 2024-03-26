package main

func main() {

	cards := newDeck()

	// cards.saveToFile("cards.txt")
	// cards := newDeckFromFile("cards.txt")

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	cards.shuffle()
	cards.print()
}
