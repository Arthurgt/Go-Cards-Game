package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)
	deckAsAstring := cards.toString()
	fmt.Println("Deck as a single string")
	fmt.Println(deckAsAstring)
	fmt.Println("Your hand")
	hand.print()
	fmt.Println("Remainig deck")
	remainingDeck.print()
	cards.saveDeckToFile("my_cards.txt")
}
