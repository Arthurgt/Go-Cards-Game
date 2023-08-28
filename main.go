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
	cardsReadedFromFile := readDeckFromFile("my_cards.txt")
	fmt.Println("Deck readed")
	cardsReadedFromFile.print()
}
