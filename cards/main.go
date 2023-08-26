package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)
	fmt.Println("Your hand")
	hand.print()
	fmt.Println("Remainig deck")
	remainingDeck.print()
}
