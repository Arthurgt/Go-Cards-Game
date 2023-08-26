package main

func main() {
	cards := deck{"Six of diamonds", newCard()}
	cards = append(cards, "Three of glances")
	cards.print()
}

func newCard() string {
	return "Five of diamonds"
}
